package cmdb

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"gorm.io/gen"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/logback"
	"github.com/vela-ssoc/vela-common-mba/netutil"
	"gorm.io/gen/field"
)

type Client interface {
	// Save 保存 cmdb 信息
	Save(ctx context.Context, dat *model.Cmdb) error

	// FetchAndSave 查询运维中心的 cmdb 信息后并保存到数据库
	FetchAndSave(ctx context.Context, id int64, inet string) error
}

func NewClient(cfg Configurer, client netutil.HTTPClient, slog logback.Logger) Client {
	return &restClient{
		cfg:    cfg,
		client: client,
		slog:   slog,
	}
}

type restClient struct {
	cfg    Configurer
	client netutil.HTTPClient
	slog   logback.Logger
	mutex  sync.RWMutex
	done   bool
	err    error
	addr   string
}

func (rc *restClient) Save(ctx context.Context, dat *model.Cmdb) error {
	if dat == nil {
		return nil
	}

	monTbl := query.Minion
	assigns := []field.AssignExpr{
		monTbl.OrgPath.Value(dat.OrgPath),
		monTbl.Identity.Value(dat.BaoleijiIdentity),
		monTbl.Category.Value(dat.Category),
		monTbl.OpDuty.Value(dat.OpDuty),
		monTbl.Comment.Value(dat.Comment),
		monTbl.IBu.Value(dat.IBu),
		monTbl.IDC.Value(dat.IDC),
	}

	info, err := monTbl.WithContext(ctx).
		Where(monTbl.ID.Eq(dat.ID), monTbl.Inet.Eq(dat.Inet)).
		UpdateColumnSimple(assigns...)
	if err != nil || info.RowsAffected == 0 {
		return err
	}

	tbl := query.Cmdb
	return tbl.WithContext(ctx).
		Where(tbl.ID.Eq(dat.ID), tbl.Inet.Eq(dat.Inet)).
		Save(dat)
}

func (rc *restClient) Save2(ctx context.Context, dat *model.Cmdb2) error {
	if dat == nil || len(dat.PrivateIP) == 0 {
		return nil
	}

	duties := make([]string, 0, 4)
	for _, m := range dat.OpDuty.Main {
		duties = append(duties, m.Username)
	}
	monTbl := query.Minion
	assigns := []field.AssignExpr{
		monTbl.Identity.Value(dat.BaoleijiIdentity),
		monTbl.OpDuty.Value(strings.Join(duties, ",")),
		monTbl.Comment.Value(dat.Comment),
		monTbl.IBu.Value(dat.Department),
		monTbl.IDC.Value(dat.IDC),
	}

	info, err := monTbl.WithContext(ctx).
		Where(monTbl.ID.Eq(dat.ID), monTbl.Inet.Eq(dat.PrivateIP[0])).
		UpdateColumnSimple(assigns...)
	if err != nil || info.RowsAffected == 0 {
		return err
	}

	tbl := query.Cmdb2
	var where gen.Condition
	switch dat.CIType {
	case "server":
		if sn := dat.SN; sn == "" {
			where = tbl.SN.Eq(sn)
		}
	case "vserver", "vmware_server":
		if uuid := dat.UUID; uuid != "" {
			where = tbl.UUID.Eq(uuid)
		}
	case "public_cloud":
		if id := dat.PublicCloudID; id != "" {
			where = tbl.PublicCloudIDC.Eq(id)
		}
	case "docker":
		if id := dat.InstanceID; id != "" {
			where = tbl.InstanceID.Eq(id)
		}
	}
	if where == nil {
		return nil
	}

	return tbl.WithContext(ctx).
		Where(where).
		Save(dat)
}

func (rc *restClient) FetchAndSave(ctx context.Context, id int64, inet string) error {
	rc.slog.Debugf("fetchAndSave cmdb: %s(%d)", inet, id)
	dat, err := rc.fetch(ctx, inet)
	if err != nil || dat == nil {
		if err != nil {
			rc.slog.Warnf("查询 cmdb 错误：%s", err)
		}
		return err
	}

	dat.ID = id
	dat.Inet = inet

	return rc.Save(ctx, dat)
}

func (rc *restClient) fetch(parent context.Context, inet string) (*model.Cmdb, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := rc.newRequest(ctx, inet)
	if err != nil {
		return nil, err
	}
	res, err := rc.client.Do(req)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	// 反序列化
	ret := new(reply)
	if err = json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	if len(ret.Result) != 0 {
		return ret.Result[0], nil
	}
	return nil, nil
}

func (rc *restClient) newRequest(ctx context.Context, inet string) (*http.Request, error) {
	addr, err := rc.cfg.Load(ctx)
	if err != nil {
		return nil, err
	}
	val := "q=private_ip:(" + inet + ")"
	if addr.RawQuery != "" {
		addr.RawQuery += "&"
	}
	addr.RawQuery += val

	return http.NewRequestWithContext(ctx, http.MethodGet, addr.String(), nil)
}
