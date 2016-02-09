package duck

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	var interfacetest interface{}
	var blah string
	blah = "65"
	interfacetest = &blah
	ptrtst := int64(5)
	cases := []struct {
		In  interface{}
		Ok  bool
		Out int64
	}{
		{"2", true, 2},
		{"2.00", true, 2},
		{"1.1", false, 0},
		{" 2 ", false, 0},
		{"Inf", false, 0},
		{"NaN", false, 0},
		{int(12), true, 12},
		{int8(12), true, 12},
		{int16(12), true, 12},
		{int32(12), true, 12},
		{int64(12), true, 12},
		{uint(12), true, 12},
		{uint8(12), true, 12},
		{uint16(12), true, 12},
		{uint32(12), true, 12},
		{uint64(12), true, 12},
		{float32(3.0), true, 3},
		{float64(3.0), true, 3},
		{3.3, false, 0},
		{math.Inf(1), false, 0},
		{math.NaN(), false, 0},
		{true, true, 1},
		{false, true, 0},
		{&ptrtst, true, 5},
		{nil, false, 0},
		{interfacetest, true, 65},
		{"false", true, 0},
		{"true", true, 1},
	}

	for _, c := range cases {
		Out, Ok := Int(c.In)
		require.Equal(t, c.Ok, Ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, Out, JSONString(c))
		}
	}
}

func TestFloat(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		In  interface{}
		Ok  bool
		Out float64
	}{
		{"2", true, 2},
		{"2.345", true, 2.345},
		{" 2 ", false, 0},
		{"Inf", true, math.Inf(1)},
		{int(12), true, 12},
		{int8(12), true, 12},
		{int16(12), true, 12},
		{int32(12), true, 12},
		{int64(12), true, 12},
		{uint(12), true, 12},
		{uint8(12), true, 12},
		{uint16(12), true, 12},
		{uint32(12), true, 12},
		{uint64(12), true, 12},
		{float32(3.0), true, 3},
		{float64(3.0), true, 3},
		{3.3, true, 3.3},
		{math.Inf(1), true, math.Inf(1)},
		{true, true, 1},
		{false, true, 0},
		{&ptrtst, true, 5},
		{nil, false, 0},
		{"false", true, 0},
		{"true", true, 1},
	}

	for _, c := range cases {
		Out, Ok := Float(c.In)
		require.Equal(t, c.Ok, Ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, Out, JSONString(c))
		}
	}

	//Testing NaN needs its own little thing
	Out, Ok := Float("NaN")
	require.True(t, Ok)
	require.True(t, math.IsNaN(Out))
	Out, Ok = Float(math.NaN())
	require.True(t, Ok)
	require.True(t, math.IsNaN(Out))

}

func TestBool(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		In  interface{}
		Ok  bool
		Out bool
	}{
		{"2", true, true},
		{"0", true, false},
		{"2.0", true, true},
		{"0.0", true, false},
		{" 2 ", false, false},
		{"Inf", true, true},
		{"NaN", true, false},
		{int(12), true, true},
		{int(0), true, false},
		{uint(12), true, true},
		{uint(0), true, false},
		{3.45, true, true},
		{0.0, true, false},
		{math.NaN(), true, false},
		{"true", true, true},
		{"false", true, false},
		{true, true, true},
		{false, true, false},
		{nil, false, false},
		{&ptrtst, true, true},
		{-1, true, false},
		{"-1", true, false},
	}

	for _, c := range cases {
		Out, Ok := Bool(c.In)
		require.Equal(t, c.Ok, Ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, Out, JSONString(c))
		}
	}
}

func TestString(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		In  interface{}
		Ok  bool
		Out string
	}{
		{"hello", true, "hello"},
		{int(12), true, "12"},
		{int(0), true, "0"},
		{uint(12), true, "12"},
		{uint(0), true, "0"},
		{3.45, true, "3.45"},
		{0.0, true, "0"},
		{3.0, true, "3"},
		{math.Inf(1), true, "+Inf"},
		{math.NaN(), true, "NaN"},
		{true, true, "true"},
		{false, true, "false"},
		{nil, false, ""},
		{&ptrtst, true, "5"},
	}

	for _, c := range cases {
		Out, Ok := String(c.In)
		require.Equal(t, c.Ok, Ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, Out, JSONString(c))
		}
	}
}

func TestJSONString(t *testing.T) {
	cases := []struct {
		In  interface{}
		Out string
	}{
		{3.0, "3"},
		{map[string]interface{}{"hi": 3}, `{"hi":3}`},
	}

	for _, c := range cases {
		require.Equal(t, c.Out, JSONString(c.In), JSONString(c))
	}
}
