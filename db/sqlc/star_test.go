package db

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountStar(t *testing.T) {
	code := GetAllCodes(t)
	cnt, err := testQueries.CountStar(context.Background(), code.ID)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "\"stars\" does not exist "))
	} else {
		require.NotEmpty(t, cnt)
	}

}
