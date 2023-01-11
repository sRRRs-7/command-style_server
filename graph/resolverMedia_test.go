package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllMedia(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			getAllMedia(limit: %d, skip: %d) {
				id
				title
				contents
				img
				created_at
				updated_at
			}
	}`, 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}
	arr, list, _ := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllMedia\"")
	require.Equal(t, true, len(list["\"id\""]) >= 3)
	require.Equal(t, true, len(list["\"title\""]) >= 3)
	require.Equal(t, true, len(list["\"contents\""]) >= 3)
	require.Equal(t, true, len(list["\"img\""]) >= 3)

	require.Zero(t, len(list["\"created_at\""]))
	require.Zero(t, len(list["\"updated_at\""]))
}

func TestGetMediaResolver(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			getMedia(id: %d) {
				id
				title
				contents
				img
				created_at
				updated_at
			}
	}`, 2)

	q := struct {
		Query string
	}{
		Query: query,
	}
	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getMedia\"")
	require.Equal(t, "\"2\"", list["\"id\""])
	require.True(t, 3 <= len(list["\"title\""]))
	require.True(t, 3 <= len(list["\"contents\""]))
	require.True(t, 3 <= len(list["\"img\""]))
	require.Contains(t, string(result), "created_at")
	require.Contains(t, string(result), "updated_at")
}

func TestUpdateMedia(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			updateMedia(id: %s, title: %s, contents: %s, img: %s) {
				is_error
				message
			}
	}`, "\"1\"", "\"title\"", "\"contents\"", "\"img\"")

	q := struct {
		Query string
	}{
		Query: query,
	}
	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"updateMedia\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateMedia OK\"", list["\"message\""])
}

func TestDeleteMedia(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			deleteMedia(id: %d) {
				is_error
				message
			}
	}`, 3)

	q := struct {
		Query string
	}{
		Query: query,
	}
	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteMedia\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteMedia OK\"", list["\"message\""])
}
