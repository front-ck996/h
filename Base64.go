package csy

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"regexp"
	"strings"
)

func ImageToBase64(inputImg image.Image) (string, error) {
	s := ""
	var buf bytes.Buffer
	b64encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	if err := imaging.Encode(b64encoder, inputImg, imaging.PNG); err != nil {
		return s, err
	}
	if err := b64encoder.Close(); err != nil {
		return s, err
	}
	encoded := buf.String()
	return fmt.Sprintf("data:image/png;base64,%s", encoded), nil

	////// 解码图像
	////img, _, err := image.Decode(file)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//
	//// 将图像编码为字节数据
	//buffer := new(bytes.Buffer)
	//err := png.Encode(buffer, inputImg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 将字节数据进行 Base64 编码
	//encoded := base64.StdEncoding.EncodeToString(buffer.Bytes())
	//
	//return fmt.Sprintf("data:image/png;base64,%s", encoded), nil

}

func Base64ToImage(base64Str string) (image.Image, error) {
	base64Str = Base64RemoveFirstTag(base64Str)
	// 解码 Base64 字符串
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	// 创建图像对象
	img, err := imaging.Decode(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}
	return img, nil

	//// 解码 Base64 字符串为字节数据
	//base64Str = Base64RemoveFirstTag(base64Str)
	//data, err := base64.StdEncoding.DecodeString(base64Str)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 创建字节数据的 Reader
	//reader := bytes.NewReader(data)
	//
	//// 解码字节数据为图像
	//img, err := png.Decode(reader)
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}
	//return img, nil
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
	matches := re.FindStringSubmatch(base64Str)

	if len(matches) > 1 {
		base64Str = matches[1]
	}

	return base64Str
}
