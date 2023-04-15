package csy

import "testing"

func TestConvCharsetToUtf8(t *testing.T) {
	content := []byte{0xa7, 0xac, 0xba, 0xc3, 0xba, 0xda, 0xd7, 0xd6, 0xbf, 0xaa, 0xd4, 0xf2, 0xce, 0xc4, 0xbd, 0xe2, 0xb2, 0xbb, 0xa3, 0xac, 0xca, 0xc7, 0xd7, 0xdc, 0xce, 0xc4, 0xa3, 0xac, 0xc7, 0xeb, 0xb2, 0xbb, 0xd2, 0xbb, 0xd6, 0xc3, 0xd3, 0xda}
	ConvCharsetToUtf8(content)

}

//func main() {
//	// 定义需要转换的字符串
//
//	// 检测输入字符串的编码格式
//	detector := chardet.NewTextDetector()
//	result, err := detector.DetectBest(content)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("Detected encoding:", result.Charset)
//
//	// 将字符串从源编码转换为 UTF-8 编码
//	decoded, err := decodeContent(content, result.Charset)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(decoded))
//}
//
//func decodeContent(content []byte, sourceEncoding string) ([]byte, error) {
//	// 获取源编码解码器
//	decoder, err := encoding.GetEncoding(sourceEncoding).NewDecoder().Bytes(content)
//	if err != nil {
//		return nil, err
//	}
//
//	// 转换为 UTF-8 编码
//	utf8Encoder := encoding.UTF8.NewEncoder()
//	reader := transform.NewReader(bytes.NewReader(decoder), utf8Encoder)
//	utf8Content, err := ioutil.ReadAll(reader)
//	if err != nil {
//		return nil, err
//	}
//	return utf8Content, nil
//}
