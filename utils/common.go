package utils

import (
	"math/rand"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
)

var hexRunes = []rune("123456789abcdef")

func RandHexStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = hexRunes[rand.Intn(len(hexRunes))]
	}
	return string(b)
}

func ParseOptions(vu modules.VU, val goja.Value) map[string]string {
	options := make(map[string]string)
	rt := vu.Runtime()

	if IsNilly(val) {
		return options
	}

	params := val.ToObject(rt)
	for _, k := range params.Keys() {
		options[k] = params.Get(k).ToString().String()
	}

	return options
}

func IsNilly(val goja.Value) bool {
	return val == nil || goja.IsNull(val) || goja.IsUndefined(val)
}
