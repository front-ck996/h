package csy

import "reflect"

func SliceUnique[T any](arr []T) (newArr []T) {
	newArr = make([]T, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if reflect.DeepEqual(arr[i], arr[j]) {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
