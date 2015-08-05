package duck

import "reflect"

//preprocessSet
func preprocessSet(i interface{}) (v reflect.Value, k reflect.Kind, ok bool) {
	v = reflect.ValueOf(i)
	k = v.Kind()
	if k != reflect.Ptr {
		return v, k, false
	}
	v = reflect.Indirect(v)
	k = v.Kind()

	//Multilevel pointers are not supported at the moment
	if k == reflect.Ptr || !v.CanSet() {
		return v, k, false
	}
	return v, k, true
}

//Set sets the interface pointer v to the given value
func Set(i, setto interface{}) (ok bool) {
	v, k, ok := preprocessSet(i)
	if !ok {
		return false
	}

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res, ok := Int(setto)
		if !ok {
			return false
		}
		v.SetInt(res)
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res, ok := Int(setto)
		if !ok || res < 0 {
			return false
		}
		v.SetUint(uint64(res))
		return true
	case reflect.Float32, reflect.Float64:
		res, ok := Float(setto)
		if !ok {
			return false
		}
		v.SetFloat(res)
		return true
	case reflect.Bool:
		res, ok := Bool(setto)
		if !ok {
			return false
		}
		v.SetBool(res)
		return true
	case reflect.String:
		res, ok := String(setto)
		if !ok {
			return false
		}
		v.SetString(res)
		return true
	case reflect.Interface:
		v.Set(reflect.ValueOf(setto))
		return true
	}
	return false
}
