package tracetest

import (
	"github.com/dop251/goja"
	"github.com/xoscar/xk6-tracetest-tracing/models"
	"go.k6.io/k6/js/modules"
)

type Tracetest struct {
	vu modules.VU
}

type NewFunc func(call goja.ConstructorCall) *goja.Object

func New(vu modules.VU) *Tracetest {
	return &Tracetest{
		vu: vu,
	}
}

func (t *Tracetest) Constructor(call goja.ConstructorCall) *goja.Object {
	rt := t.vu.Runtime()
	isCliInstalled := t.getIsCliInstalled()

	if !isCliInstalled {
		panic("The tracetest CLI is not installed. Please install it before using this module")
	}

	return rt.ToValue(t).ToObject(rt)
}

func (t *Tracetest) RunTest(testID, traceID string) (*models.TracetestRun, error) {
	definitionFileName, err := t.exportTest(testID)
	if err != nil {
		return nil, err
	}

	run, err := t.runTest(definitionFileName, traceID)
	return run, err
}

func (t *Tracetest) RunFromDefinition(testDefinition, traceID string) (*models.TracetestRun, error) {
	run, err := t.runFromDefinition(testDefinition, traceID)

	return run, err
}
