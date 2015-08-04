package duck

import (
	"math"
	"reflect"
)

//Variables used for Cmp
var (
	LessThan    = -1
	GreaterThan = 1
	Equals      = 0
	CantCompare = -2
)

//Cmp performs a comparison between the two values, and returns the result
//of the comparison (LessThan, GreaterThan, Equals, CantCompare), which are defined as ints
func Cmp(arg1 interface{}, arg2 interface{}) int {
	eq, ok := Equal(arg1, arg2)
	if !ok {
		return CantCompare
	}
	if eq {
		return Equals
	}

	lt, _ := Lt(arg1, arg2)
	if lt {
		return LessThan
	}

	f1, _ := Float(arg1)
	f2, _ := Float(arg2)
	if math.IsNaN(f1) || math.IsNaN(f2) {
		return CantCompare
	}
	return GreaterThan
}

//Lt returns true if arg1 < arg2
func Lt(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	//In order to compare the two, use the lowest common denominator: float
	f1, ok := Float(arg1)
	if !ok {
		return false, false
	}
	f2, ok := Float(arg2)

	return f1 < f2, ok
}

//Lte returns true if arg1 <= arg2
func Lte(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	f1, ok := Float(arg1)
	if !ok {
		return false, false
	}
	f2, ok := Float(arg2)
	return f1 <= f2, ok
}

//Gt returns true if arg1 > arg2
func Gt(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	return Lt(arg2, arg1)
}

//Gte returns true if arg1 >= arg2
func Gte(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	return Lte(arg2, arg1)
}

//Equal attempts to check equality between two interfaces. If the values
//are not directly comparable thru DeepEqual, tries to do a "duck" comparison.
//One thing to note: Equal(x,y) != Equal(y,x) in some cases! This is because the first
//argument to Equal is special - it is the "benchmark" type - the thing we want to be
//comparing arg2 against
//	true true -> true
//	"true" true -> true
//	"1" true -> true
//	1.0 1 -> true
//	1.345 "1.345" -> true
//	50.0 true -> true
//	0.0 false -> true
func Equal(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	if reflect.DeepEqual(arg1, arg2) {
		return true, true
	}

	//The values are different - let's see if we can create a valid comparison

	_, k1 := preprocess(arg1)
	_, k2 := preprocess(arg2)

	if k1 == k2 && k1 != reflect.String {
		//The kinds are the same - DeepEqual should have handled it - it is false!
		//EXCEPT for when it is string - two strings, "2" and "2.0" have the same meaning
		//but are not equal
		return false, true
	}

	//TODO: There is the special case of comparing a char with a string

	//Now attempt to compare equality float-wise
	f1, ok := Float(arg1)
	if !ok {
		return false, false
	}
	f2, ok := Float(arg2)

	if math.IsNaN(f1) && math.IsNaN(f2) {
		return true, ok
	}

	return f1 == f2, ok

}

//Eq is short-hand for Equal. Look at Equal for detailed description
func Eq(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	return Equal(arg1, arg2)
}
