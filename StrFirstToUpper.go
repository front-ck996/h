package csy

import "strings"

func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func StrFirstToLower(str string) string {
	runStr := []rune(str)
	return strings.ToLower(string(runStr[:1])) + string(runStr[1:])

	//temp := strings.Split(str, "_")
	//var upperStr string
	//for y := 0; y < len(temp); y++ {
	//	vv := []rune(temp[y])
	//	if y != 0 {
	//		for i := 0; i < len(vv); i++ {
	//			if i == 0 {
	//				vv[i] -= 32
	//				upperStr += string(vv[i]) // + string(vv[i+1])
	//			} else {
	//				upperStr += string(vv[i])
	//			}
	//		}
	//	}
	//}
	//return strings.ToLower(temp[0]) + upperStr
}
