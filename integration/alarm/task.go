package alarm

import (
	"context"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/integration/devops"
	"github.com/vela-ssoc/vela-common-mb/integration/dong"
	"github.com/vela-ssoc/vela-common-mb/integration/formwork"
	"github.com/vela-ssoc/vela-common-mb/logback"
)

type sendTask struct {
	dat  any
	sub  *model.Subscriber
	rend formwork.Render
	slog logback.Logger
	dong dong.Client
	dps  devops.Client
}

func (st *sendTask) Run() {
	// 发送告警
	if dongs := st.sub.Dong; len(dongs) != 0 {
		st.sendDong(dongs)
	}
	if devs := st.sub.Devops; len(devs) != 0 {
		st.sendDevops(devs)
	}
}

func (st *sendTask) sendDong(dongs []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	title, body := st.rend.RiskDong(ctx, st.dat, st.dat)
	err := st.dong.Send(ctx, dongs, nil, title, body)
	st.slog.Infof("发送风险 %s 结果：%v", dongs, err)
}

func (st *sendTask) sendDevops(devs []*model.Devops) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// title, body := st.rend.RiskDong(ctx, st.dat, st.dat)
	err := st.dps.Send(ctx, "告警", "内容", devs)
	st.slog.Infof("发送风险 %s 结果：%v", devs, err)
}
