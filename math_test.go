package duck

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, true, 3},
		{"hello", "world", true, "helloworld"},
		{132, 1.0, true, 133},
		{12, -1, true, 11},
		{nil, 4, true, 4},
	}

	for _, c := range cases {
		out, ok := Add(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.EqualValues(t, c.Out, out, JSONString(c))
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, true, 1},
		{"hello", "world", false, nil},
		{132, 1.0, true, 131},
		{12, -1, true, 13},
	}

	for _, c := range cases {
		out, ok := Subtract(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.EqualValues(t, c.Out, out, JSONString(c))
		}
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, true, 2},
		{"hello", "world", false, nil},
		{3, 2.0, true, 6},
	}

	for _, c := range cases {
		out, ok := Multiply(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.EqualValues(t, c.Out, out, JSONString(c))
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, true, 2},
		{"hello", "world", false, nil},
		{3, 2.0, true, 1.5},
		{3, 0, true, math.Inf(1)},
	}

	for _, c := range cases {
		out, ok := Divide(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.EqualValues(t, c.Out, out, JSONString(c))
		}
	}
}

func TestMod(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, true, 0},
		{"hello", "world", false, nil},
		{3, 2.0, true, 1.0},
	}

	for _, c := range cases {
		out, ok := Mod(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.EqualValues(t, c.Out, out, JSONString(c))
		}
	}
}

func BenchmarkAddInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add(23, 47)
	}
}

func BenchmarkAddFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add(23.3, 47.2)
	}
}

func BenchmarkAddString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add("hello world!", " hi")
	}
}

func BenchmarkMultiplyFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lt(23.4, 12.3)
	}
}

func BenchmarkModFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lt(23.4, 12.3)
	}
}
