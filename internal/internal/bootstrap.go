package internal

import (
	"github.com/hetue/core/internal/internal/internal"
	"github.com/hetue/core/internal/internal/param"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	param *param.Bootstrap
}

func NewBootstrap(param *param.Bootstrap) *Bootstrap {
	return &Bootstrap{
		param: param,
	}
}

func (b *Bootstrap) Boot() {
	application := pangu.New()
	if "" != b.param.Name {
		application.Name(b.param.Name)
	}
	if "" != b.param.Usage {
		application.Usage(b.param.Usage)
	}
	if "" != b.param.Copyright {
		application.Copyright(b.param.Copyright)
	}
	for key, value := range b.param.Metadata {
		application.Metadata(key, value)
	}
	application.Get().Run(internal.NewBootstrap)
}
