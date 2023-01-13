package graph

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	CreateToken(t)
}

func CreateToken(t *testing.T) string {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createToken(username: %s)
	}`, "\"srrrs\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://127.0.0.1:8080/query", "")

	arr := strings.Split(string(result), ":")
	token := string(arr[2])[1 : len(arr[2])-3]

	require.Equal(t, list["\"data\""], "\"createToken\"")

	return token
}

func TestAdminCreateToken(t *testing.T) {
	CreateToken(t)
}

func CreateAdminToken(t *testing.T) string {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createAdminToken(username: %s, password: %s)
	}`, "\"srrrs\"", "\"srrrs\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://127.0.0.1:8080/admin/query", "")

	arr := strings.Split(string(result), ":")
	token := string(arr[2])[1 : len(arr[2])-3]

	require.Equal(t, list["\"data\""], "\"createAdminToken\"")
	return token
}
