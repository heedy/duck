package duck

import "reflect"

//valueGetDuckTag gets the element of the struct with the given duck tag
//it is assumed that v is a value of a struct.
func valueGetDuckTag(v reflect.Value, tag string) (val reflect.Value, ok bool) {
	//As of now, I don't think it is possible to do this without looping through the fields
	//of the struct. Since structs have a constant number of fields, this shouldn't be an issue
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		val := t.Field(i)
		if val.Tag.Get("duck") == tag {
			return v.FieldByName(val.Name), true
		}
	}
	return v, false
}

//valueGetStep runs one step of getting a subobject from an object
func valueGetStep(v reflect.Value, elem interface{}) (val reflect.Value, ok bool) {
	k := v.Kind()
	for k == reflect.Ptr || k == reflect.Interface {
		if k == reflect.Interface {
			v = v.Elem()
		} else {
			v = reflect.Indirect(v)
		}
		k = v.Kind()
	}

	switch k {
	case reflect.Map:
		estr, ok := String(elem)
		if !ok {
			return v, false
		}

		//Only string map[string] are supported right now...
		rv := v.MapIndex(reflect.ValueOf(estr))
		if rv.IsValid() {
			return rv, true
		}

		return rv, false

	case reflect.Array, reflect.Slice, reflect.String:
		//Get the element of the array requested (assuming that elem is an index)
		eint, ok := Int(elem)
		if !ok {
			return v, false
		}
		if v.Len() <= int(eint) {
			return v, false
		}

		//Allow python-like indexing for arrays
		if eint < 0 {
			if int(eint) < -v.Len() {
				return v, false
			}
			return v.Index(v.Len() + int(eint)), true
		}

		return v.Index(int(eint)), true

	case reflect.Struct:
		//Get the element as a string, so that we can see if the struct has such an element
		estr, ok := String(elem)
		if !ok {
			return v, false
		}

		t := v.Type()

		val, ok := t.FieldByName(estr)
		if !ok {
			return valueGetDuckTag(v, estr)
		}
		//If the duck tag is not there or it is the same, return the interface
		if val.Tag.Get("duck") == "" || val.Tag.Get("duck") == estr {
			return v.FieldByName(estr), true
		}

		return valueGetDuckTag(v, estr)
	}
	return v, false
}

//valueGet takes an interface object, and the list of elements of sub-objects which
//are to be traversed. It returns the value after traversing through all of the elements
func valueGet(v reflect.Value, elem ...interface{}) (val reflect.Value, ok bool) {
	for e := range elem {
		v, ok = valueGetStep(v, elem[e])
		if !ok {
			return v, false
		}
	}
	return v, true
}

//Get takes an object, and the element name/index to extract, and returns the element, as well as an
//ok boolean specifying if the element was found. Remember that only exported fields are available from structs
//One note: only `map[string]` is supported right now. maps of another type will cause get to panic!
//This is an active weakness in the current implementation. It should be enough to get data from arbitrary marshalled json.
//
//Get also supports multiple args for multilevel data:
//	//returns deeplyNestedStruct->foo->bar
//	duck.Get(deeplyNestedStruct, "foo","bar")
func Get(i interface{}, elem ...interface{}) (val interface{}, ok bool) {
	v, ok := valueGet(reflect.ValueOf(i), elem...)
	if ok {
		return v.Interface(), ok
	}
	return nil, false
}

//Set attempts to set a subelement of the passed-in object. In the same way
//that Get gets the value, Set allows to set it. The only possibly confusing
//thing about Set is that the "to" value and the element names are reversed
//to allow for multiple arguments (as in Get)
//
//	v := map[string]string{"hi":"world"}
//	duck.Set(&v,45,"hi")
//	//now v["hi"] == "45"
//
//	v2 := map[string]interface{}{"foo":map[string]interface{}{"bar": 23}}
//	duck.Set(&v2,"hello!","foo","bar")
//	//v2["foo"]["bar"] == "hello!"
//
//NOTE: Currently setting structs (ie, if setto is a struct) is non-functional
//if the receiving end is not interface{}. Also, only map[string]interface{} is supported
//of maps to receive values - and Set will only replace existing values.
func Set(i interface{}, setto interface{}, elem ...interface{}) (ok bool) {

	v := reflect.ValueOf(i)

	k := v.Kind()
	if k == reflect.Ptr {
		v = v.Elem()
	}

	//maps are weird in that their values cannot just be set (CanSet=False).
	//we therefore have to check if the last value is a map or not
	var lastelement interface{}
	if len(elem) > 0 {
		lastelement = elem[len(elem)-1]
		elem = elem[:len(elem)-1]
	}

	v, ok = valueGet(v, elem...)
	if !ok {
		return false
	}

	if v.Kind() == reflect.Map && lastelement != nil {
		//We are to set a map value...
		estr, ok := String(lastelement)
		vestr := reflect.ValueOf(estr)
		if !ok {
			return false
		}

		rv := v.MapIndex(vestr)
		if !rv.IsValid() {
			return false
		}

		//Alright, so we get the type information from the value
		switch rv.Kind() {
		case reflect.Interface:
			v.SetMapIndex(vestr, reflect.ValueOf(setto))
			return true
			/*TODO: Support more map types
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				res, ok := Int(setto)
				if !ok {
					return false
				}
			*/

		}
		return false
	} else if lastelement != nil {
		v, ok = valueGet(v, lastelement)
	}
	if !ok || !v.CanSet() {
		return false
	}

	switch v.Kind() {
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
