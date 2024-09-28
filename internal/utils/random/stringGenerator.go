package random

import "github.com/thanhpk/randstr"

func StringGenerator(length int) string {
	token := randstr.Hex(length)
	return token
}
