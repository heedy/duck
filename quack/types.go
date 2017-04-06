package quack

import (
	"encoding/json"
	"strconv"

	"github.com/connectordb/duck/quack/fconv"
)

// Int converts the given interface into an int
// if possible.
func Int(i interface{}) (int64, bool) {
	switch i := i.(type) {
	case int64:
		return i, true
	case int:
		return int64(i), true
	case float64:
		iv := int64(i)
		return iv, float64(iv) == i
	case string:
		fv, ok := fconv.ParseFloat(i)
		iv := int64(fv)
		return iv, ok && float64(iv) == fv
	case float32:
		iv := int64(i)
		return iv, float32(iv) == i
	case int8:
		return int64(i), true
	case int16:
		return int64(i), true
	case int32:
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint64:
		return int64(i), true
	case uint:
		return int64(i), true
	case bool:
		if i {
			return 1, true
		}
		return 0, true
	case nil:
		return 0, true
	default:
		return 0, false
	}
}

// Float converts the given interface into an int
// if possible.
func Float(i interface{}) (float64, bool) {
	switch i := i.(type) {
	case float64:
		return i, true
	case float32:
		return float64(i), true
	case int64:
		return float64(i), true
	case int:
		return float64(i), true
	case string:
		return fconv.ParseFloat(i)
	case int8:
		return float64(i), true
	case int16:
		return float64(i), true
	case int32:
		return float64(i), true
	case uint8:
		return float64(i), true
	case uint16:
		return float64(i), true
	case uint32:
		return float64(i), true
	case uint64:
		return float64(i), true
	case uint:
		return float64(i), true
	case bool:
		if i {
			return 1, true
		}
		return 0, true
	case nil:
		return 0, true
	default:
		return 0, false
	}
}

// Bool converts the given interface into a boolean
// if possible.
func Bool(i interface{}) (bool, bool) {
	switch i := i.(type) {
	case bool:
		return i, true
	case float64:
		return i > 0, true
	case float32:
		return i > 0, true
	case int64:
		return i > 0, true
	case int:
		return i > 0, true
	case string:
		// ParseFloat is modified to make 'true' == 1 and 'false' == 0
		f, ok := fconv.ParseFloat(i)
		return f > 0, ok
	case int8:
		return i > 0, true
	case int16:
		return i > 0, true
	case int32:
		return i > 0, true
	case uint8:
		return i > 0, true
	case uint16:
		return i > 0, true
	case uint32:
		return i > 0, true
	case uint64:
		return i > 0, true
	case uint:
		return i > 0, true
	case nil:
		return false, true
	default:
		return false, false
	}
}

//String attempts to convert the given interface into a string. Examples:
//	1 -> "1"
//	2.45 -> "2.45"
//	"hi" -> "hi"
//	false -> "false"
//	true -> "true"
// Bool converts the given interface into a boolean
// if possible.
func String(i interface{}) (string, bool) {
	switch i := i.(type) {
	case string:
		return i, true
	case bool:
		if i {
			return "true", true
		}
		return "false", true
	case float64:
		return strconv.FormatFloat(i, 'g', -1, 64), true
	case float32:
		return strconv.FormatFloat(float64(i), 'g', -1, 32), true
	case int64:
		return strconv.FormatInt(i, 10), true
	case int:
		return strconv.FormatInt(int64(i), 10), true
	case int8:
		return strconv.FormatInt(int64(i), 10), true
	case int16:
		return strconv.FormatInt(int64(i), 10), true
	case int32:
		return strconv.FormatInt(int64(i), 10), true
	case uint8:
		return strconv.FormatUint(uint64(i), 10), true
	case uint16:
		return strconv.FormatUint(uint64(i), 10), true
	case uint32:
		return strconv.FormatUint(uint64(i), 10), true
	case uint64:
		return strconv.FormatUint(uint64(i), 10), true
	case uint:
		return strconv.FormatUint(uint64(i), 10), true
	case nil:
		return "", true
	default:
		return "", false
	}
}

//JSONString attempts to convert to a string... but if that fails it
//marshals the given data into json, and returns the json string. If it
//can't marshal, it just returns an empty string
func JSONString(i interface{}) string {
	s, ok := String(i)
	if ok {
		return s
	}
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b)
}
