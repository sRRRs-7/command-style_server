package db

import (
	"context"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestListMedia(t *testing.T) {
	arg := ListMediaParams{
		Limit:  30,
		Offset: 0,
	}
	medias, err := testQueries.ListMedia(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(medias) >= 0)
}

func TestCreateMedia(t *testing.T) {
	title := utils.RandomString(10)
	contents := utils.RandomString(50)

	arg := CreateMediaParams{
		Title:     title,
		Contents:  contents,
		Img:       []byte{10},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := testQueries.CreateMedia(context.Background(), arg)
	require.NoError(t, err)
}

func TestGetMedia(t *testing.T) {
	media, err := testQueries.GetMedia(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, media)
}

func TestUpdateMedia(t *testing.T) {
	arg := UpdateMediaParams{
		ID:       1,
		Title:    "hello",
		Contents: "fine",
		Img:      []byte{10},
	}
	err := testQueries.UpdateMedia(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteMedia(t *testing.T) {
	err := testQueries.DeleteMedia(context.Background(), 1)
	require.NoError(t, err)
}
