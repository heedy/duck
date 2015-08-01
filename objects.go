package duck

import "reflect"

//Get takes an object, and the element name/index to extract, and returns the element, as well as an
//ok boolean specifying if the element was found. Remember that only exported fields are available from structs
//One note: only `map[string]` is supported right now. maps of another type will cause get to panic!
//This is an active weakness in the current implementation. It should be enough to get data from arbitrary marshalled json.
func Get(i interface{}, elem interface{}) (val interface{}, ok bool) {
	v, k := preprocess(i)

	switch k {
	case reflect.Map:
		estr, ok := String(elem)
		if !ok {
			return nil, false
		}

		//Only string map[string] are supported right now...
		rv := v.MapIndex(reflect.ValueOf(estr))
		if rv.IsValid() {
			return rv.Interface(), true
		}

		return nil, false
	case reflect.Array, reflect.Slice, reflect.String:
		//Get the element of the array requested (assuming that elem is an index)
		eint, ok := Int(elem)
		if !ok {
			return nil, false
		}
		if v.Len() <= int(eint) {
			return nil, false
		}

		//Allow python-like indexing for arrays
		if eint < 0 {
			if int(eint) < -v.Len() {
				return nil, false
			}
			return v.Index(v.Len() + int(eint)).Interface(), true
		}

		return v.Index(int(eint)).Interface(), true

	case reflect.Struct:
		//Get the element as a string, so that we can see if the struct has such an element
		estr, ok := String(elem)
		if !ok {
			return nil, false
		}

		val := v.FieldByName(estr)
		if val.IsValid() {
			return val.Interface(), true
		}
		return nil, false
	}
	return nil, false
}
