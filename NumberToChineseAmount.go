package csy

import "strconv"

var chineseDigits = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
var chineseUnits = []string{"", "拾", "佰", "仟", "万", "拾万", "佰万", "仟万", "亿", "拾亿", "佰亿", "仟亿"}

func NumberToChineseAmount(number float64) string {
	intPart := int(number)
	fracPart := int((number - float64(intPart)) * 100)

	intPartStr := strconv.Itoa(intPart)
	fracPartStr := strconv.Itoa(fracPart)

	result := ""

	// 处理整数部分
	for i := 0; i < len(intPartStr); i++ {
		digit := int(intPartStr[i] - '0')
		if digit != 0 {
			result += chineseDigits[digit] + chineseUnits[len(intPartStr)-i-1]
		} else {
			if i < len(intPartStr)-1 && int(intPartStr[i+1]-'0') != 0 {
				result += chineseDigits[digit]
			}
		}
	}

	// 处理小数部分
	if fracPart != 0 {
		result += "点"
		for i := 0; i < len(fracPartStr); i++ {
			digit := int(fracPartStr[i] - '0')
			result += chineseDigits[digit]
		}
	}

	return result
}
