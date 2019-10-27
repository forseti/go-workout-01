package c12reflect

import "reflect"

func walk(x interface{}, fn func(input string))  {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}



	//numberOfValues := 0
	//var getField func(int) reflect.Value
	//
	//switch val.Kind() {
	//case reflect.String:
	//	fn(val.String())
	//case reflect.Struct:
	//	numberOfValues = val.NumField()
	//	getField = val.Field
	//case reflect.Slice, reflect.Array:
	//	numberOfValues = val.Len()
	//	getField = val.Index
	//case reflect.Map:
	//	for _, key := range val.MapKeys() {
	//		walk(val.MapIndex(key).Interface(), fn)
	//	}
	//}
	//
	//for i := 0; i < numberOfValues; i++ {
	//	walk(getField(i).Interface(), fn)
	//}


	//fn("I still can't believe South Korea beat Germany to put them last in their group")

	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		walk(val.Index(i).Interface(), fn)
	//	}
	//
	//	return
	//}
	//
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//
	//	//if field.Kind() == reflect.String {
	//	//	fn(field.String())
	//	//}
	//	//
	//	//if field.Kind() == reflect.Struct {
	//	//	walk(field.Interface(), fn)
	//	//}
	//
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		walk(field.Interface(), fn)
	//	}
	//}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
