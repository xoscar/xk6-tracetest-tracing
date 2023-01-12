package utils

import "math/rand"

var hexRunes = []rune("123456789abcdef")

func RandHexStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = hexRunes[rand.Intn(len(hexRunes))]
	}
	return string(b)
}
