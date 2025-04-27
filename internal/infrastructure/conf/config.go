package conf

import (
	"time"

	pconf "github.com/pinguo-icc/kratos-library/v2/conf"
	pmongo "github.com/pinguo-icc/kratos-library/v2/mongo"
	"github.com/pinguo-icc/kratos-library/v2/trace"
)

type Bootstrap struct {
	App     *App
	Http    *HTTP
	Grpc    *GRPC
	Trace   *trace.Config
	MongoDB *pmongo.Config
}

func Load(env string) (*Bootstrap, error) {
	out := new(Bootstrap)
	err := pconf.Load(env, out, nil)
	return out, err
}

type (
	HTTP struct {
		Address string
		Timeout time.Duration
	}

	GRPC struct {
		Address string
		Timeout time.Duration
	}
	App struct {
		Name   string
		Env    string
		Region string // 部署位置： hz(杭州)、sg(新加坡)
	}
)
