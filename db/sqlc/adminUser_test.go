package db

import (
	"context"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestAdminUser(t *testing.T) {
	username := utils.RandomUser(5)
	password := utils.RandomPassword(8)

	// create admin user
	arg := CreateAdminUserParams{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := testQueries.CreateAdminUser(context.Background(), arg)
	require.NoError(t, err)

	// get admin user
	arg2 := GetAdminUserParams{
		Username: username,
		Password: password,
	}
	user, err := testQueries.GetAdminUser(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, username)
	require.Equal(t, user.Password, password)

	require.NotZero(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)
}
