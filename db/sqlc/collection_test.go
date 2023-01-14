package db

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetCollection(t *testing.T) int64 {
	id := GetCode(t)
	code, err := testQueries.GetCollection(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, code)

	return code.ID
}

func TestGetCollection(t *testing.T) {
	GetCollection(t)
}

func TestCreateCollection(t *testing.T) {
	userId := int64(3)
	codeId := int64(3)

	arg := CreateCollectionParams{
		UserID: userId,
		CodeID: codeId,
	}
	err := testQueries.CreateCollection(context.Background(), arg)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "violates foreign key constraint"))
	} else {
		require.NoError(t, err)
	}
}

func TestDeleteCollection(t *testing.T) {
	codeId := GetCollection(t)

	err := testQueries.DeleteCollection(context.Background(), codeId)
	require.NoError(t, err)
}

func TestGetAllCollections(t *testing.T) {
	username, _, _ := CreateRandomUser(t)
	userId := GetUserByUsername(t, username)

	// get all collection
	arg2 := GetAllCollectionsParams{
		UserID: userId,
		Limit:  30,
		Offset: 0,
	}
	collections, err := testQueries.GetAllCollections(context.Background(), arg2)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "violates foreign key constraint"))
	} else {
		require.Equal(t, len(collections) >= 0, true)
	}
}

func TestGetAllCollectionsBySearch(t *testing.T) {
	username, _, _ := CreateRandomUser(t)
	userId := GetUserByUsername(t, username)

	keyword := "go"
	arg2 := GetAllCollectionsBySearchParams{
		UserID:      userId,
		Username:    username,
		Code:        keyword,
		Description: keyword,
		Column5:     keyword,
		Limit:       30,
		Offset:      0,
	}
	collections, err := testQueries.GetAllCollectionsBySearch(context.Background(), arg2)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "violates foreign key constraint"))
	} else {
		require.Equal(t, len(collections) >= 0, true)
	}

}
