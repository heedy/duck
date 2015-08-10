package duck

import "math"

//Add attempts to add two interfaces. It works similarly to expectations everywhere, with
//the extra 'extreme string ducktyping' twist.
//	"hello" + "world" = "helloworld"
//	true + 1 = 2
//	3.14 + 1 = 4.14
//	"23"+"42" = 65
func Add(i1, i2 interface{}) (res interface{}, ok bool) {
	f1, ok := Float(i1)
	f2, ok2 := Float(i2)

	if ok && ok2 {
		return f1 + f2, true
	}

	//At least one of the two is not convertible to float. See if they are
	//both strings
	i1s, ok := i1.(string)
	i2s, ok2 := i2.(string)
	if ok && ok2 {
		return i1s + i2s, true
	}
	return nil, false
}

//Subtract attempts to perform i1-i2. It does not have any special weirdness, it simply
//converts the two to floats, and subtracts them.
func Subtract(i1, i2 interface{}) (res float64, ok bool) {
	f1, ok := Float(i1)
	f2, ok2 := Float(i2)

	if ok && ok2 {
		return f1 - f2, true
	}
	return 0, false
}

//Multiply tries to convert the two to numbers and subtract them
func Multiply(i1, i2 interface{}) (res float64, ok bool) {
	f1, ok := Float(i1)
	f2, ok2 := Float(i2)

	if ok && ok2 {
		return f1 * f2, true
	}
	return 0, false
}

//Divide attempts to convert to numbers, and divide the first by the second
func Divide(i1, i2 interface{}) (res float64, ok bool) {
	f1, ok := Float(i1)
	f2, ok2 := Float(i2)

	if ok && ok2 {
		return f1 / f2, true
	}
	return 0, false
}

//Mod finds the i1%i2.
func Mod(i1, i2 interface{}) (res float64, ok bool) {
	f1, ok := Float(i1)
	f2, ok2 := Float(i2)

	if ok && ok2 {
		return math.Mod(f1, f2), true
	}
	return 0, false
}
