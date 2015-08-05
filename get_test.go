package duck

import (
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
