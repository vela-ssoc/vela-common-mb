package sonatype

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mba/netutil"
)

// Searcher 漏洞查询
type Searcher interface {
	// Search 通过 purl 查询漏洞库
	Search([]string) ([]*model.SBOMVuln, error)
}

func NewClient(cfg Configurer, cli netutil.HTTPClient) Searcher {
	return &sonaClient{
		cfg:   cfg,
		cli:   cli,
		batch: 100, // sonatype 每次最大只能查询 120 条 purl
	}
}

type sonaClient struct {
	cfg   Configurer
	cli   netutil.HTTPClient
	batch int
}

func (sc *sonaClient) Search(purls []string) ([]*model.SBOMVuln, error) {
	if len(purls) == 0 {
		return nil, nil
	}
	cts, err := sc.find(purls)
	if err != nil {
		return nil, err
	}

	// 转换层
	ret := cts.Vulns()

	return ret, nil
}

func (sc *sonaClient) find(purls []string) (sonatypeComponents, error) {
	size := len(purls)
	if size <= sc.batch {
		return sc.query(purls)
	}

	sb := new(sonatypeBatch)

	return sb.finds(sc, purls)
}

func (sc *sonaClient) query(purls []string) (sonatypeComponents, error) {
	// 读取配置项
	addr, auth, err := sc.cfg.Load(context.Background())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	header := http.Header{"Authorization": []string{auth}}
	body := &sonatypeRequest{Coordinates: purls}
	var reply sonatypeComponents
	err = sc.cli.JSON(ctx, http.MethodPost, addr, body, &reply, header)

	return reply, err
}

// sonatypeBatch 分批次向 sonatype 查询漏洞库
type sonatypeBatch struct {
	wg     sync.WaitGroup
	mutex  sync.Mutex
	err    error
	result sonatypeComponents
}

func (sb *sonatypeBatch) finds(st *sonaClient, purls []string) (sonatypeComponents, error) {
	batch := st.batch
	for size := len(purls); size > 0; {
		end := batch
		if end > size {
			end = size
		}

		shards := purls[:end]
		purls = purls[end:]
		sb.wg.Add(1)

		go sb.find(st, shards)
	}

	sb.wg.Wait()

	return sb.result, sb.err
}

func (sb *sonatypeBatch) find(st *sonaClient, purls []string) {
	defer sb.wg.Done()

	// 联网查询漏洞库
	components, err := st.query(purls)

	sb.mutex.Lock()
	defer sb.mutex.Unlock()

	if err != nil {
		sb.err = err
	} else {
		sb.result = append(sb.result, components...)
	}
}
