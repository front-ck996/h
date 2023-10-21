package csy

import (
	"crypto/md5"
	"io"
)

func Md5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return string(h.Sum(nil))
}
