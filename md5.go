package csy

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	//h := md5.New()
	//io.WriteString(h, s)
	data := []byte(s) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
