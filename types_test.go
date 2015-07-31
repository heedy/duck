package duck

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		in  interface{}
		ok  bool
		out int64
	}{
		{"2", true, 2},
		{"2.00", true, 2},
		{"1.1", false, 0},
		{" 2 ", false, 0},
		{"Inf", false, 0},
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
	}

	for _, c := range cases {
		out, ok := Int(c.in)
		require.Equal(t, c.ok, ok, fmt.Sprintf("%v", c))
		if c.ok {
			require.Equal(t, c.out, out, fmt.Sprintf("%v", c))
		}
	}

}

func TestFloat(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		in  interface{}
		ok  bool
		out float64
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
	}

	for _, c := range cases {
		out, ok := Float(c.in)
		require.Equal(t, c.ok, ok, fmt.Sprintf("%v", c))
		if c.ok {
			require.Equal(t, c.out, out, fmt.Sprintf("%v", c))
		}
	}

	//Testing NaN needs its own little thing
	out, ok := Float("NaN")
	require.True(t, ok)
	require.True(t, math.IsNaN(out))
	out, ok = Float(math.NaN())
	require.True(t, ok)
	require.True(t, math.IsNaN(out))

}

func TestBool(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		in  interface{}
		ok  bool
		out bool
	}{
		{"2", true, true},
		{"0", true, false},
		{"2.0", true, true},
		{"0.0", true, false},
		{" 2 ", false, false},
		{"Inf", true, true},
		{int(12), true, true},
		{int(0), true, false},
		{uint(12), true, true},
		{uint(0), true, false},
		{3.45, true, true},
		{0.0, true, false},
		{"true", true, true},
		{"false", true, false},
		{true, true, true},
		{false, true, false},
		{nil, false, false},
		{&ptrtst, true, true},
	}

	for _, c := range cases {
		out, ok := Bool(c.in)
		require.Equal(t, c.ok, ok, fmt.Sprintf("%v", c))
		if c.ok {
			require.Equal(t, c.out, out, fmt.Sprintf("%v", c))
		}
	}
}

func TestString(t *testing.T) {
	ptrtst := int64(5)
	cases := []struct {
		in  interface{}
		ok  bool
		out string
	}{
		{"hello", true, "hello"},
		{int(12), true, "12"},
		{int(0), true, "0"},
		{uint(12), true, "12"},
		{uint(0), true, "0"},
		{3.45, true, "3.45"},
		{0.0, true, "0"},
		{3.0, true, "3"},
		{true, true, "true"},
		{false, true, "false"},
		{nil, false, ""},
		{&ptrtst, true, "5"},
	}

	for _, c := range cases {
		out, ok := String(c.in)
		require.Equal(t, c.ok, ok, fmt.Sprintf("%v", c))
		if c.ok {
			require.Equal(t, c.out, out, fmt.Sprintf("%v", c))
		}
	}
}

func TestJSONString(t *testing.T) {
	cases := []struct {
		in  interface{}
		ok  bool
		out string
	}{
		{3.0, true, "3"},
		{map[string]interface{}{"hi": 3}, true, `{"hi":3}`},
	}

	for _, c := range cases {
		out, ok := JSONString(c.in)
		require.Equal(t, c.ok, ok, fmt.Sprintf("%v", c))
		if c.ok {
			require.Equal(t, c.out, out, fmt.Sprintf("%v", c))
		}
	}
}
