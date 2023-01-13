package httpClient

import (
	"context"
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/xoscar/xk6-tracetest-tracing/models"
	"github.com/xoscar/xk6-tracetest-tracing/utils"
	"go.k6.io/k6/js/modules"
	k6HTTP "go.k6.io/k6/js/modules/k6/http"
)

var defaultPropagatorList = []models.PropagatorName{
	models.PropagatorW3C,
}

func getOptions(vu modules.VU, val goja.Value) (Options, error) {
	rawOptions := utils.ParseOptions(vu, val)
	options := Options{
		Propagator: models.NewPropagator(defaultPropagatorList),
	}

	if len(rawOptions) == 0 {
		return options, nil
	}

	for key, value := range rawOptions {
		switch key {
		case "propagator":
			rawPropagatorList := strings.Split(value, ",")
			propagatorList := make([]models.PropagatorName, len(rawPropagatorList))
			for i, propagator := range rawPropagatorList {
				propagatorList[i] = models.PropagatorName(propagator)
			}

			options.Propagator = models.NewPropagator(propagatorList)
			//TODO: validate
		default:
			return options, fmt.Errorf("unknown HTTP tracing option '%s'", key)
		}
	}

	return options, nil
}

func requestToHttpFunc(method string, request HttpRequestFunc) HttpFunc {
	return func(ctx context.Context, url goja.Value, args ...goja.Value) (*k6HTTP.Response, error) {
		return request(method, url, args...)
	}
}
