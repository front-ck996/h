package csy

import (
	"net/http"
	"os"
)

func GetFileContentType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 仅嗅探内容类型的第一个
	// 使用了 512 个字节。

	buf := make([]byte, 512)
	_, err = file.Read(buf)

	if err != nil {
		return "", err
	}

	// 真正起作用的函数
	contentType := http.DetectContentType(buf)

	return contentType, nil
}
