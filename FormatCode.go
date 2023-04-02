package csy

import (
	"go/format"
)

func FormatCode(src []byte) []byte {
	got, _ := format.Source(src)
	return got
}
