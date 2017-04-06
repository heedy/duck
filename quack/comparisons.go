package quack

import "math"

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

	f1, ok := Float(arg1)
	if !ok {
		return CantCompare
	}
	f2, ok := Float(arg2)
	if !ok {
		return CantCompare
	}

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
//	true true -> true
//	"true" true -> true
//	"1" true -> true
//	1.0 1 -> true
//	1.345 "1.345" -> true
//	50.0 true -> true
//	0.0 false -> true
func Equal(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	if arg1 == arg2 {
		return true, true
	}
	// Comparing things to nil should work
	if arg1 == nil || arg2 == nil {
		// The first if statement should have been equal if they are both not nil
		return false, true
	}

	s1, ok1 := arg1.(string)
	s2, ok2 := arg2.(string)
	if ok1 && ok2 {
		if s1 == s2 {
			return true, true
		}
	}

	// Neither is a string. Let's see if one of them is a number.
	// That way we know we can compare as numbers
	f1, ok := Float(arg1)
	if !ok {
		// using this slows doewn the entire function. It is the same issue
		// that required forking parseFloat to be faster.
		//if reflect.DeepEqual(arg1, arg2) {
		//	return true, true
		//}
		return false, true
	}

	// OK, so at least the first one was a number. Let's try the second one
	f2, ok := Float(arg2)
	if !ok {
		return false, true
	}

	if math.IsNaN(f1) && math.IsNaN(f2) {
		return true, true
	}

	return f1 == f2, true

}

//Eq is short-hand for Equal. Look at Equal for detailed description
func Eq(arg1 interface{}, arg2 interface{}) (res bool, ok bool) {
	return Equal(arg1, arg2)
}
