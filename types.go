package duck

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func preprocess(i interface{}) (reflect.Value, reflect.Kind) {
	v := reflect.ValueOf(i)
	k := v.Kind()
	for k == reflect.Ptr {
		v = reflect.Indirect(v)
		k = v.Kind()
	}
	return v, k
}

//Int attempts to convert the given interface into an integer
//	"1" -> 1
//	1 -> 1
//	1.0 -> 1
//	"1.0" -> 1
//	false -> 0
//	true -> 1
//	"1.3" -> !ok
//	nil -> !ok
func Int(i interface{}) (res int64, ok bool) {
	v, k := preprocess(i)

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		res = int64(f)
		if float64(res) == f {
			return res, true
		}
		return res, false
	case reflect.Bool:
		if v.Bool() {
			return 1, true
		}
		return 0, true
	case reflect.String:
		f, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, false
		}
		res = int64(f)
		if float64(res) == f {
			return res, true
		}
		return res, false
	}

	return 0, false
}

//Float attempts to convert the given interface into a float
//	1 -> 1.0
//	"1.34" -> 1.34
//	1.34 -> 1.34
//	false -> 0.0
//	true -> 1.0
//	" 2.3 " -> !ok
func Float(i interface{}) (res float64, ok bool) {
	v, k := preprocess(i)

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true

	case reflect.Bool:
		if v.Bool() {
			return 1, true
		}
		return 0, true
	case reflect.String:
		f, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, false
		}
		return f, true
	}

	return 0, false
}

//Bool attempts to convert the given interface into a boolean.
//There can be a loss of information in this one.
//	"hi" -> !ok
//	true -> true
//	"true" -> true
//	"false" -> false
//	"0.0" -> false
//	"3.34" -> true
//	1 -> true
//	1337 -> true
// 1.434 -> true
// -1 -> true
//	nil -> false
//	0.0 -> false
//	0 -> false
//	"" -> false
func Bool(i interface{}) (res bool, ok bool) {
	v, k := preprocess(i)

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0, true
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0, true
	case reflect.Bool:
		return v.Bool(), true
	case reflect.String:
		s := v.String()
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			if s == "false" || len(s) == 0 {
				return false, true
			} else if s == "true" {
				return true, true
			}
			return false, false
		}
		return f != 0.0, true
	}

	return false, false
}

//String attempts to convert the given interface into a string. Examples:
//	1 -> "1"
//	2.45 -> "2.45"
//	"hi" -> "hi"
//	false -> "false"
//	true -> "true"
func String(i interface{}) (res string, ok bool) {
	v, k := preprocess(i)

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), true
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64), true
	case reflect.Bool:
		if v.Bool() {
			return "true", true
		}
		return "false", true
	case reflect.String:
		return v.String(), true

	}

	return "", false
}

//JSONString attempts to convert to a string... but if that fails it
//marshals the given data into json, and returns the json string
func JSONString(i interface{}) (res string, ok bool) {
	s, ok := String(i)
	if ok {
		return s, ok
	}
	b, err := json.Marshal(i)
	if err != nil {
		return "", false
	}
	return string(b), true
}
