package tracetest

import (
	"github.com/xoscar/xk6-tracetest-tracing/modules/instance"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/tracetest", New())
}

type (
	RootModule struct{}
)

var _ modules.Module = &RootModule{}

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return instance.New(vu)
}
