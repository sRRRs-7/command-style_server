package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMedia(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createMedia(title: %s, contents: %s, img: %s) {
				is_error
				message
			}
	}`, "\"title\"", "\"contents\"", "\"img\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)

	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createMedia\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateMedia OK\"", list["\"message\""])
}

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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")

	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllMedia\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetMediaResolver(t *testing.T) {
	NewServer()

	id := 3
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
	}`, id)

	q := struct {
		Query string
	}{
		Query: query,
	}
	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")

	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getMedia\"")
	require.Equal(t, fmt.Sprintf("\"%d\"", id), list["\"id\""])
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

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query", "")

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
	arr, list, result := NewRequest(t, q, "http://localhost:8080/query", "")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteMedia\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteMedia OK\"", list["\"message\""])
}
