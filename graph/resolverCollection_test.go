package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// fix token session for redis
func TestCreateAdminCollection(t *testing.T) {
	NewServer()

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

	_, list, result := NewRequest(t, q, "http://localhost:8080/admin/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createAdminCollection\"")
	require.Equal(t, true, len(list["\"is_error\""]))
	require.Equal(t, true, len(list["\"message\""]))
}

// fix cookie session
func TestCreateCollection(t *testing.T) {
	NewServer()

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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createAdminCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateAdminCollection OK\"", list["\"message\""])
}

// fix cookie session
func TestGetCollection(t *testing.T) {
	NewServer()

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
		}`, 7)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateCollection OK\"", list["\"message\""])
}

func TestDeleteCollection(t *testing.T) {
	NewServer()

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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteCollection\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteCollection OK\"", list["\"message\""])
}

// fix cookie session
func TestGetAllCollection(t *testing.T) {
	NewServer()

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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteCollection\"")
	require.True(t, len(string(result)) >= 1000)
}

// fix cookie session
func TestGetAllCollectionBySearch(t *testing.T) {
	NewServer()

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
		}`, "\"gogogogo\"", 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"deleteCollection\"")
	require.True(t, len(string(result)) <= 100)
}
