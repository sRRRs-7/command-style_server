package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckRegexp(t *testing.T) {
	reg, s := "\\S", "Go"
	require.True(t, CheckRegex(reg, s))

	reg, s = "\\s", " "
	require.Equal(t, CheckRegex(reg, s), true)

	reg, s = "[A-Za-z0-9]+", "1sdd"
	require.Equal(t, CheckRegex(reg, s), true)
}

func TestRegexpArray(t *testing.T) {
	reg, s := "[:,]", "a,b:c:d,e"

	arr := RegexpArray(reg, s)
	require.Equal(t, 5, len(arr))
}

func TestRetrieveRegexp(t *testing.T) {
	reg, s := "\".*\"", "abc\"hello\"hello"

	str := RetrieveRegexp(reg, s)
	require.Equal(t, "\"hello\"", str)
}
