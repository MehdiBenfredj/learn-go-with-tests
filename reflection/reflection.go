package reflection

import (
	"reflect"
)

func walk(x any, fn func(input string)) {
	fields := extractFields(x)
	for _, field := range fields {
		fn(field)
	}
}

func extractFields(x any) []string {
	stringFields := []string{}
	val := reflect.ValueOf(x)
	numberFields := val.NumField()
	for numField := 0; numField < numberFields; numField++ {
		var field = val.Field(numField)
		if field.Kind() == reflect.String {
			stringFields = append(stringFields, field.String())
		} else if field.Kind() == reflect.Struct {
			stringFields = append(stringFields, extractFields(field.Interface())...)
		}
	}
	return stringFields
}

/* field
Profile struct {
	Age  int
	City string
}
*/
