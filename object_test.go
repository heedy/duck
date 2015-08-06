package duck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	A1 string
	A2 int
}

type TestStruct2 struct {
	A1 int `duck:"-"`
	A2 int
	A3 int `duck:"lol"`
}

func TestGet(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"23", 0, true, byte('2')},
		{"23", -1, true, byte('3')},
		{"23", 50, false, nil},
		{"23", -50, false, nil},
		{"23", "blah", false, nil},
		{[]int{1, 2, 3, 4, 5}, "1.0", true, int(2)},
		{TestStruct{"foo", 4}, "A1", true, "foo"},
		{TestStruct{"foo", 4}, nil, false, ""},
		{TestStruct{"foo", 4}, "notexist", false, ""},
		{TestStruct2{1, 2, 3}, "A1", false, 0},
		{TestStruct2{1, 2, 3}, "A3", false, 0},
		{TestStruct2{1, 2, 3}, "A2", true, 2},
		{TestStruct2{1, 2, 3}, "lol", true, 3},
		{34.5, 1, false, nil},
		{map[string]interface{}{"true": 5}, true, true, 5},
		{map[string]interface{}{"true": 5}, "true", true, 5},
		{map[string]interface{}{"true": 5}, "zoo", false, 5},
		{map[string]interface{}{"true": 5}, nil, false, 5},
	}

	for _, c := range cases {
		out, ok := Get(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, out, JSONString(c))
		}

	}
}

func TestMultlevelGet(t *testing.T) {
	d := map[string]interface{}{"foo": map[string]int{"bar": 1337}}
	out, ok := Get(d, "foo", "bar")
	require.True(t, ok)
	require.Equal(t, 1337, out)

	out, ok = Get(d, "foo", "nofoo")
	require.False(t, ok)
}

func setTest(t *testing.T, val interface{}, setto interface{}, vals []interface{}, isok bool) {
	ok := Set(val, setto, vals...)
	if !isok {
		require.False(t, ok, fmt.Sprintf("%v->%v", val, setto))
		return
	}
	require.True(t, ok, fmt.Sprintf("%v->%v", val, setto))

}

func TestSet(t *testing.T) {
	var caseInterface interface{}
	var caseInt int
	var caseUint uint
	var caseString string
	var caseFloat float64
	var caseBool bool

	//Test successful set
	setTest(t, &caseInt, "4", nil, true)
	require.Equal(t, 4, caseInt)
	setTest(t, &caseUint, "4", nil, true)
	require.Equal(t, uint(4), caseUint)
	setTest(t, &caseFloat, 3, nil, true)
	require.Equal(t, float64(3.0), caseFloat)
	setTest(t, &caseBool, 1, nil, true)
	require.Equal(t, true, caseBool)
	setTest(t, &caseString, "hi", nil, true)
	require.Equal(t, "hi", caseString)
	setTest(t, &caseInterface, 54, nil, true)
	require.Equal(t, 54, caseInterface)

	//Test fail set
	setTest(t, caseInt, "4", nil, false) //Need ptr

	setTest(t, &caseInt, "lol", nil, false)
	setTest(t, &caseUint, "lol", nil, false)
	setTest(t, &caseFloat, "lol", nil, false)
	setTest(t, &caseBool, "lol", nil, false)
	setTest(t, &caseString, map[string]string{"hi": "hi"}, nil, false)

	s := TestStruct{"foo", 135}
	setTest(t, &s, 24, []interface{}{"A2"}, true)
	require.Equal(t, 24, s.A2)

	arr := []int{1, 2, 3, 4}
	setTest(t, &arr, 24, []interface{}{"1"}, true)
	require.Equal(t, 24, arr[1])

	g := map[string]interface{}{"hi": "world"}
	setTest(t, &g, 1337, []interface{}{"noexists"}, false)
	setTest(t, &g, 1337, []interface{}{"hi"}, true)
	require.Equal(t, 1337, g["hi"])

	g2 := map[string]string{"hi": "world"}
	setTest(t, &g2, 1337, []interface{}{"hi"}, false)

	deepnest := []interface{}{&s}
	setTest(t, &deepnest, 49, []interface{}{0, "A2"}, true)
	require.Equal(t, 49, s.A2)
}

func TestLength(t *testing.T) {
	i, ok := Length([]int{1, 2, 3})
	require.True(t, ok)
	require.Equal(t, 3, i)
	i, ok = Length("hello")
	require.True(t, ok)
	require.Equal(t, 5, i)

	i, ok = Length(5)
	require.False(t, ok)
}
