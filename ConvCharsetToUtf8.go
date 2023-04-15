package csy

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/saintfish/chardet"
)

func ConvCharsetToUtf8(content []byte) (string, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(content)
	if err != nil {
		return "", err
	}
	fmt.Println(content)
	fmt.Println("Detected encoding:", result.Charset)
	// 将字符串从源编码转换为 UTF-8 编码
	// 获取源编码解码器
	//charset := result.Charset
	//if result.Language == "zh" {
	//	charset = "gbk"
	//}

	encoder := mahonia.NewEncoder("gbk")
	convertString := encoder.ConvertString(string(content))
	fmt.Println("convertString", convertString)
	return convertString, nil
}
