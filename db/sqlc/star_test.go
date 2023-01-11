package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountStar(t *testing.T) {
	code := GetAllCodes(t)
	cnt, err := testQueries.CountStar(context.Background(), code.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cnt)
}
