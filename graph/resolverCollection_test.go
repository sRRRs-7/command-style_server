package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAdminCollection(t *testing.T) {
	query := fmt.Sprintf(`
		mutation {
			createAdminCollection(user_id: %d, code_id: %d) {
				is_error
				message
			}
		}`, 7, 6)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateAdminToken(t)
	_, list, result := NewAdminRequest(t, q, "http://localhost:8080/admin/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createAdminCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateAdminCollection OK\"", list["\"message\""])
}

func TestCreateCollection(t *testing.T) {
	query := fmt.Sprintf(`
		mutation {
			createCollection(code_id: %d) {
				is_error
				message
			}
		}`, 7)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateCollection OK\"", list["\"message\""])
}

func TestGetCollection(t *testing.T) {
	id := 7

	query := fmt.Sprintf(`
		mutation {
			getCollection(id: %d) {
				id
				username
				code
				img
				description
				performance
				star
				tags
				created_at
				updated_at
				access
				user_id
			}
		}`, id)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getCollection\"")
	require.Equal(t, fmt.Sprintf("\"%d\"", id), list["\"id\""])
	require.True(t, 3 <= len(list["\"username\""]))
	require.True(t, 3 <= len(list["\"code\""]))
	require.True(t, 3 <= len(list["\"img\""]))
	require.True(t, 3 <= len(list["\"description\""]))
	require.True(t, 3 <= len(list["\"performance\""]) || 0 <= len(list["\"performance\""]))
	require.Contains(t, string(result), "star")
	require.Contains(t, string(result), "tags")
	require.Contains(t, string(result), "created_at")
	require.Contains(t, string(result), "updated_at")
	require.Contains(t, string(result), "access")
	require.Contains(t, string(result), "user_id")
}

func TestDeleteCollection(t *testing.T) {
	query := fmt.Sprintf(`
		mutation {
			deleteCollection(id: %d) {
				is_error
				message
			}
		}`, 10)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteCollection OK\"", list["\"message\""])
}

func TestGetAllCollection(t *testing.T) {
	query := fmt.Sprintf(`
		query {
			getAllCollection(limit: %d, skip: %d) {
				id
				username
				code
				img
				description
				performance
				star
				tags
				created_at
				updated_at
				access
				collection_id
				user_id
			}
		}`, 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllCollection\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetAllCollectionBySearch(t *testing.T) {
	query := fmt.Sprintf(`
		query {
			getAllCollectionBySearch(keyword: %s, limit: %d, skip: %d) {
				id
				username
				code
				img
				description
				performance
				star
				tags
				created_at
				updated_at
				access
				collection_id
				user_id
			}
		}`, "\"gogo\"", 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllCollectionBySearch\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}
