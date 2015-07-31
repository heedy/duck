package duck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntOK(t *testing.T) {
	i, ok := Int("2")
	require.True(t, ok)
	require.Equal(t, int64(2), i)

	i, ok = Int("2.00")
	require.True(t, ok)
	require.Equal(t, int64(2), i)

	i, ok = Int(int8(12))
	require.True(t, ok)
	require.Equal(t, int64(12), i)

	i, ok = Int(int16(12))
	require.True(t, ok)
	require.Equal(t, int64(12), i)

	i, ok = Int(int32(12))
	require.True(t, ok)
	require.Equal(t, int64(12), i)

	i, ok = Int(int64(12))
	require.True(t, ok)
	require.Equal(t, int64(12), i)

	i, ok = Int(float32(3.0))
	require.True(t, ok)
	require.Equal(t, int64(3), i)

	i, ok = Int(float64(3.0))
	require.True(t, ok)
	require.Equal(t, int64(3), i)

	i, ok = Int(false)
	require.True(t, ok)
	require.Equal(t, int64(0), i)

	i, ok = Int(true)
	require.True(t, ok)
	require.Equal(t, int64(1), i)

	ptrtst := int64(5)
	i, ok = Int(&ptrtst)
	require.True(t, ok)
	require.Equal(t, int64(5), i)
}

func TestIntNOK(t *testing.T) {
	_, ok := Int(" 2 ")
	require.False(t, ok)

	_, ok = Int(nil)
	require.False(t, ok)

	_, ok = Int("1.1")
	require.False(t, ok)

	_, ok = Int(1.1)
	require.False(t, ok)
}

func TestFloatOK(t *testing.T) {
	i, ok := Float("2")
	require.True(t, ok)
	require.Equal(t, float64(2), i)

	i, ok = Float("2.345")
	require.True(t, ok)
	require.Equal(t, float64(2.345), i)

	i, ok = Float(int8(12))
	require.True(t, ok)
	require.Equal(t, float64(12), i)

	i, ok = Float(int16(12))
	require.True(t, ok)
	require.Equal(t, float64(12), i)

	i, ok = Float(int32(12))
	require.True(t, ok)
	require.Equal(t, float64(12), i)

	i, ok = Float(int64(12))
	require.True(t, ok)
	require.Equal(t, float64(12), i)

	i, ok = Float(float32(3.0))
	require.True(t, ok)
	require.Equal(t, float64(3), i)

	i, ok = Float(float64(3.0))
	require.True(t, ok)
	require.Equal(t, float64(3), i)

	i, ok = Float(false)
	require.True(t, ok)
	require.Equal(t, float64(0), i)

	i, ok = Float(true)
	require.True(t, ok)
	require.Equal(t, float64(1), i)

	ptrtst := int64(5)
	i, ok = Float(&ptrtst)
	require.True(t, ok)
	require.Equal(t, float64(5), i)
}

func TestFloatNOK(t *testing.T) {
	_, ok := Float(" 2 ")
	require.False(t, ok)

	_, ok = Float(nil)
	require.False(t, ok)

}

func TestBoolOK(t *testing.T) {
	i, ok := Bool("2")
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool("0")
	require.True(t, ok)
	require.False(t, i)

	i, ok = Bool("2.0")
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool("0.0")
	require.True(t, ok)
	require.False(t, i)

	i, ok = Bool("true")
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool("false")
	require.True(t, ok)
	require.False(t, i)

	i, ok = Bool(2)
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool(0)
	require.True(t, ok)
	require.False(t, i)

	i, ok = Bool(1.0)
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool(0.0)
	require.True(t, ok)
	require.False(t, i)

	i, ok = Bool(true)
	require.True(t, ok)
	require.True(t, i)

	i, ok = Bool(false)
	require.True(t, ok)
	require.False(t, i)
}

func TestBoolNOK(t *testing.T) {
	_, ok := Bool(" 2 ")
	require.False(t, ok)

	_, ok = Bool(nil)
	require.False(t, ok)

}

func TestStringOK(t *testing.T) {
	i, ok := String("hello")
	require.True(t, ok)
	require.Equal(t, "hello", i)

	i, ok = String(int8(12))
	require.True(t, ok)
	require.Equal(t, "12", i)

	i, ok = String(int16(12))
	require.True(t, ok)
	require.Equal(t, "12", i)

	i, ok = String(int32(12))
	require.True(t, ok)
	require.Equal(t, "12", i)

	i, ok = String(int64(12))
	require.True(t, ok)
	require.Equal(t, "12", i)

	i, ok = String(float32(3.0))
	require.True(t, ok)
	require.Equal(t, "3", i)

	i, ok = String(float64(3.0))
	require.True(t, ok)
	require.Equal(t, "3", i)

	i, ok = String(false)
	require.True(t, ok)
	require.Equal(t, "false", i)

	i, ok = String(true)
	require.True(t, ok)
	require.Equal(t, "true", i)

	ptrtst := int64(5)
	i, ok = String(&ptrtst)
	require.True(t, ok)
	require.Equal(t, "5", i)
}

func TestStringNOK(t *testing.T) {
	_, ok := String(nil)
	require.False(t, ok)

}

func TestJSONString(t *testing.T) {
	i, ok := JSONString(float32(3.0))
	require.True(t, ok)
	require.Equal(t, "3", i)

	d := map[string]interface{}{"hi": 3}
	i, ok = JSONString(d)
	require.True(t, ok)
	require.Equal(t, `{"hi":3}`, i)
}
