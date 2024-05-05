package csy

import (
	"go/format"
)

func FormatCode(src []byte) ([]byte, error) {
	got, err := format.Source(src)
	if len(src) == 0 {
		return src, nil
	}
	if len(got) == 0 {
		return src, nil
	}

	return got, err
}

// ReplaceStrToXinXin 字符串固定长度转 *
func ReplaceStrToXinXin(str string, start, end int) (result string) {
	runes := []rune(str)
	var sfz []rune
	if len(runes) > start {
		sfz = append(sfz, runes[0:start]...)
	}
	if len(runes) > start+end {
		size := len(runes) - start + end
		d := ""
		for i := 0; i < size; i++ {
			d += "*"
		}
		sfz = append(sfz, []rune(d)...)
	}
	if len(str) > start+end {
		sfz = append(sfz, runes[len(runes)-2:len(runes)]...)
	}

	if len(sfz) != 0 {
		return string(sfz)
	} else {
		return str
	}
}
