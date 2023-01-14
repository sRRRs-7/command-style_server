package db

import (
	"context"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func ListMedia(t *testing.T) *Media {
	arg := ListMediaParams{
		Limit:  30,
		Offset: 0,
	}
	medias, err := testQueries.ListMedia(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(medias) >= 0)

	return medias[0]
}

func TestListMedia(t *testing.T) {
	ListMedia(t)
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
	m := ListMedia(t)
	media, err := testQueries.GetMedia(context.Background(), m.ID)
	require.NoError(t, err)
	require.NotEmpty(t, media)
}

func TestUpdateMedia(t *testing.T) {
	m := ListMedia(t)

	arg := UpdateMediaParams{
		ID:       m.ID,
		Title:    "hello",
		Contents: "fine",
		Img:      []byte{10},
	}
	err := testQueries.UpdateMedia(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteMedia(t *testing.T) {
	m := ListMedia(t)
	err := testQueries.DeleteMedia(context.Background(), m.ID)
	require.NoError(t, err)
}
