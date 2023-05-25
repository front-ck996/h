package csy

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"regexp"
	"strings"
)

func ImageToBase64(inputImg image.Image) (string, error) {

	//// 解码图像
	//img, _, err := image.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 将图像编码为字节数据
	buffer := new(bytes.Buffer)
	err := jpeg.Encode(buffer, inputImg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 将字节数据进行 Base64 编码
	encoded := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return fmt.Sprintf("data:image/png;base64,%s", encoded), nil

}

func BaseToImage(base64Str string) (image.Image, error) {
	// 解码 Base64 字符串为字节数据
	data, err := base64.StdEncoding.DecodeString(Base64RemoveFirstTag(base64Str))
	if err != nil {
		return nil, err
	}

	// 创建字节数据的 Reader
	reader := bytes.NewReader(data)

	// 解码字节数据为图像
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return img, nil
}

func Base64Encode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

func Base64Decode(str string) ([]byte, error) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}

func Base64RemoveFirstTag(base64Str string) string {
	// 使用正则表达式提取 Base64 字符串部分
	re := regexp.MustCompile(`^data:image\/\w+;base64,(.+)`)
	return re.ReplaceAllString(base64Str, "")
}
