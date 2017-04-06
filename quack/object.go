package quack

func getIndex(elem interface{}, length int64) (int64, bool) {
	i, ok := Int(elem)
	if !ok {
		return 0, false
	}
	if length <= i {
		return 0, false
	}
	if i < 0 {
		i = length + i
		if i < 0 {
			return 0, false
		}
	}
	return i, true

}

// Get is a bit different from duck's get, since it only supports one level,
// and it only supports the quick-types
func Get(elem interface{}, t interface{}) (val interface{}, ok bool) {
	switch t := t.(type) {
	case string:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i : i+1], true
	case []interface{}:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i], true
	case map[string]interface{}:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok
	case []string:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i], true
	case []float64:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i], true
	case []int64:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i], true
	case []int:
		i, ok := getIndex(elem, int64(len(t)))
		if !ok {
			return nil, false
		}
		return t[i], true
	case map[string]string:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok
	case map[string]float64:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok
	case map[string]int64:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok

	case map[string]int:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok
	case map[string]bool:
		s, ok := String(elem)
		if !ok {
			return nil, false
		}
		v, ok := t[s]
		return v, ok

	default:
		return nil, false
	}
}

//Length extracts the length of an array/map (anything that can have len() run on it)
func Length(i interface{}) (l int, ok bool) {
	switch i := i.(type) {
	case string:
		return len(i), true
	case []interface{}:
		return len(i), true
	case map[string]interface{}:
		return len(i), true
	case []int64:
		return len(i), true
	case []float64:
		return len(i), true
	case []bool:
		return len(i), true
	case map[string]float64:
		return len(i), true
	case map[string]string:
		return len(i), true
	case map[string]bool:
		return len(i), true
	default:
		return 0, false
	}
}
