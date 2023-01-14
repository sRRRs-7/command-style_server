package db

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) (username, password, email string) {
	username = utils.RandomUser(5)
	password = utils.RandomPassword(8)
	email = utils.RandomEmail()

	arg := CreateUserParams{
		Username:    username,
		Password:    password,
		Email:       email,
		Sex:         "man",
		DataOfBirth: "1996/08/25",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	return username, password, email
}

func GetUserByUsername(t *testing.T, username string) int64 {
	user, err := testQueries.GetUserByUsername(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user.ID
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByUsername(t *testing.T) {
	username, _, _ := CreateRandomUser(t)
	id := GetUserByUsername(t, username)
	require.NotEmpty(t, id)
}

func TestGetUserByID(t *testing.T) {
	username, password, email := CreateRandomUser(t)

	id := GetUserByUsername(t, username)

	user, err := testQueries.GetUserByID(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, username)
	require.Equal(t, user.Password, password)
	require.Equal(t, user.Email, email)
	require.Equal(t, user.Sex, "man")
	require.Equal(t, user.DataOfBirth, "1996/08/25")

	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.UpdatedAt)

	require.NotZero(t, user.ID)

}

func TestDeleteUser(t *testing.T) {
	id := int64(1)
	err := testQueries.DeleteUser(context.Background(), id)
	require.NoError(t, err)
}

func TestLoginUser(t *testing.T) {
	username, password, email := CreateRandomUser(t)

	arg := LoginUserParams{
		Username: username,
		Password: password,
	}
	user, err := testQueries.LoginUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, username)
	require.Equal(t, user.Password, password)
	require.Equal(t, user.Email, email)
	require.Equal(t, user.Sex, "man")
	require.Equal(t, user.DataOfBirth, "1996/08/25")

	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.UpdatedAt)

	require.NotZero(t, user.ID)
}

func TestUpdateUser(t *testing.T) {
	username, _, email := CreateRandomUser(t)
	mail := utils.RandomEmail()

	arg := UpdateUserParams{
		Username:   username,
		Username_2: "john",
		Email:      mail,
		UpdatedAt:  time.Now(),
	}
	err := testQueries.UpdateUser(context.Background(), arg)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "users_username_key"))
	} else {
		require.NotEqual(t, username, arg.Username_2)
		require.NotEqual(t, email, mail)
	}
}
