package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAdminUser(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createAdminUser(username: %s, password: %s) {
				is_error
				message
			}
	}`, "\"srrrs\"", "\"srrrs\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateAdminToken(t)
	_, _, result := NewAdminRequest(t, q, "http://localhost:8080/admin/query", token)

	fmt.Println(string(result))

	require.Contains(t, string(result), "adminuser_username_key")
	require.Contains(t, string(result), "duplicate key value violates unique constraint")
}

func TestGetAdminUser(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
	mutation {
		getAdminUser(username: %s, password: %s) {
			is_username
			is_password
		}
	}`, "\"srrrs\"", "\"srrrs\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewAdminRequest(t, q, "http://localhost:8080/query", "")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAdminUser\"")
	require.Contains(t, string(result), "is_username")
	require.Contains(t, string(result), "true")
	require.Contains(t, string(result), "is_password")
	require.Contains(t, string(result), "true")

}
