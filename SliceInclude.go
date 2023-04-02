package csy

import "reflect"

func SliceInclude[T any](src []T, target T) bool {
	for _, element := range src {
		if reflect.DeepEqual(target, element) {
			return true
		}
	}
	return false
}
