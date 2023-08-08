package csy

import "reflect"

func IsError(v interface{}) bool {
	_, ok := v.(error)
	return ok
}
func IsErrorReflect(v interface{}) bool {
	errorType := reflect.TypeOf((*error)(nil)).Elem()
	valueType := reflect.TypeOf(errorType)
	return valueType.Implements(errorType)
}
