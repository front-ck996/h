package csy

import (
	"go/format"
)

func FormatCode(src []byte) ([]byte, error) {
	got, err := format.Source(src)
	return got, err
}
