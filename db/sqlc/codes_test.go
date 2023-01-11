package db

import (
	"context"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomCode(t *testing.T) string {
	username, _, _ := CreateRandomUser(t)
	code := utils.RandomString(10)
	description := utils.RandomString(20)
	access := int64(1)

	// create code
	arg1 := CreateCodeParams{
		Username:    username,
		Code:        code,
		Img:         []byte{10},
		Description: description,
		Performance: "",
		Star:        []int64{1, 2},
		Tags:        []string{"go"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Access:      access,
	}
	err := testQueries.CreateCode(context.Background(), arg1)
	require.NoError(t, err)

	return username
}

func GetCode(t *testing.T) int64 {
	id := int64(3)
	code, err := testQueries.GetCode(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, code)

	return code.ID
}

func GetAllCodes(t *testing.T) *Codes {
	arg := GetAllCodesParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, codes)
	require.Equal(t, len(codes) >= 1, true)

	return codes[0]
}
func TestCodes(t *testing.T) {
	CreateRandomCode(t)
}

func TestGetCode(t *testing.T) {
	GetCode(t)
}

func TestGetAllCodes(t *testing.T) {
	GetAllCodes(t)
}

func TestGetAllCodesByKeyword(t *testing.T) {
	keyword := "search"
	arg := GetAllCodesByKeywordParams{
		Username:    keyword,
		Code:        keyword,
		Description: keyword,
		Limit:       30,
		Offset:      0,
	}
	codes, err := testQueries.GetAllCodesByKeyword(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetCodesByTag(t *testing.T) {
	tags := make([]string, 10)
	tags[0] = "go"
	arg := GetAllCodesByTagParams{
		Column1:  tags[0],
		Column2:  tags[1],
		Column3:  tags[2],
		Column4:  tags[3],
		Column5:  tags[4],
		Column6:  tags[5],
		Column7:  tags[6],
		Column8:  tags[7],
		Column9:  tags[8],
		Column10: tags[9],
		Limit:    30,
		Offset:   0,
	}
	codes, err := testQueries.GetAllCodesByTag(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllCodesSortAccess(t *testing.T) {
	arg := GetAllCodesSortedAccessParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodesSortedAccess(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllCodesSortStar(t *testing.T) {
	arg := GetAllCodesSortedStarParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodesSortedStar(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllOwnCodes(t *testing.T) {
	username := CreateRandomCode(t)
	arg := GetAllOwnCodesParams{
		Username: username,
		Limit:    30,
		Offset:   0,
	}
	codes, err := testQueries.GetAllOwnCodes(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestUpdateAccess(t *testing.T) {
	arg := UpdateAccessParams{
		ID:     3,
		Access: 1,
	}
	err := testQueries.UpdateAccess(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateStar(t *testing.T) {
	arg := UpdateStarParams{
		ID:   3,
		Star: []int64{1, 2, 3},
	}
	err := testQueries.UpdateStar(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateCode(t *testing.T) {
	code := utils.RandomString(10)
	description := utils.RandomString(20)
	arg := UpdateCodeParams{
		ID:          3,
		Code:        code,
		Img:         []byte{10},
		Description: description,
		Performance: "",
		Tags:        []string{"go"},
		UpdatedAt:   time.Now(),
	}
	err := testQueries.UpdateCode(context.Background(), arg)
	require.NoError(t, err)
}
