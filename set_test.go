package duck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func setTest(t *testing.T, val interface{}, setto interface{}, isok bool) {
	ok := Set(val, setto)
	if !isok {
		require.False(t, ok)
		return
	}
	require.True(t, ok)

}

func TestSet(t *testing.T) {
	var caseInterface interface{}
	var caseInt int
	var caseUint uint
	var caseString string
	var caseFloat float64
	var caseBool bool

	//Test successful set
	setTest(t, &caseInt, "4", true)
	require.Equal(t, 4, caseInt)
	setTest(t, &caseUint, "4", true)
	require.Equal(t, uint(4), caseUint)
	setTest(t, &caseFloat, 3, true)
	require.Equal(t, float64(3.0), caseFloat)
	setTest(t, &caseBool, 1, true)
	require.Equal(t, true, caseBool)
	setTest(t, &caseString, "hi", true)
	require.Equal(t, "hi", caseString)
	setTest(t, &caseInterface, 54, true)
	require.Equal(t, 54, caseInterface)

	//Test fail set
	setTest(t, caseInt, "4", false) //Need ptr

	setTest(t, &caseInt, "lol", false)
	setTest(t, &caseUint, "lol", false)
	setTest(t, &caseFloat, "lol", false)
	setTest(t, &caseBool, "lol", false)
	setTest(t, &caseString, map[string]string{"hi": "hi"}, false)
}
