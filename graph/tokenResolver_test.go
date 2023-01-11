package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	CreateToken(t)
}

func CreateToken(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createToken(username: %s)
	}`, "\"vvvsadavvvv\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, _ := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createToken\"")
}

func TestAdminCreateToken(t *testing.T) {
	CreateToken(t)
}

func CreateAdminToken(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createAdminToken(username: %s, password: %s)
	}`, "\"vvvsadavvvv\"", "\"vvvsadavvvv\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, _ := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createAdminToken\"")
}
