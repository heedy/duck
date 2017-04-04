package duck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	cases := []struct {
		In  interface{}
		Ok  bool
		Out interface{}
	}{
		{"2", true, "2"},
		{2, true, float64(2)},
		{map[string]interface{}{"hi": "ho"}, true, map[string]interface{}{"hi": "ho"}},
	}

	for _, c := range cases {
		Out, err := Copy(c.In)
		if c.Ok {
			require.NoError(t, err, JSONString(c))
			require.Equal(t, c.Out, Out, JSONString(c))
		} else {
			require.Error(t, err)
		}
	}
}

func BenchmarkMapCopy(b *testing.B) {
	v := map[string]interface{}{
		"abc": "def",
		"efg": 234,
		"qre": 345.435,
	}

	for n := 0; n < b.N; n++ {
		Copy(v)
	}
}
