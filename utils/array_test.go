package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStarContains(t *testing.T) {
	tags := []int64{2, 3, 4, 5, 1, 8, 9}
	userId := int64(1)
	arr := StarContains(tags, userId)
	require.Equal(t, arr, []int64{2, 3, 4, 5, 9, 8})

	tags = []int64{2, 3, 4, 5, 8, 9}
	userId = int64(1)
	arr = StarContains(tags, userId)
	require.Equal(t, arr, []int64{2, 3, 4, 5, 8, 9, 1})

	tags = []int64{}
	userId = int64(1)
	arr = StarContains(tags, userId)
	require.Equal(t, arr, []int64{1})
}

func TestCreateMap(t *testing.T) {
	list := make(map[string]string)

	key1, value1 := "key1", "value1"
	list = CreateMap(key1, value1, list)

	key2, value2 := "key2", "value2"
	list = CreateMap(key2, value2, list)

	require.Equal(t, list[key1], value1)
	require.Equal(t, list[key2], value2)
}
