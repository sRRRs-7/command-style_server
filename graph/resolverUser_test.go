package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserResolver(t *testing.T) {
	NewServer()

	// hashPass, err := hash.HashPassword("ssssssss")
	// if err != nil {
	// 	t.Fatal("failed to password hash: ", err)
	// }

	query := fmt.Sprintf(`
		mutation {
			createUser(username: %s, password: %s, email: %s, sex: %s, date_of_birth: %s) {
				is_error
				message
			}
	}`, "\"vvvsadavvvv\"", "\"vvvvdsadvvvvvv\"", "\"abcef@asdsc.com\"", "\"man\"", "\"1996-13-45\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createUser\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateUser OK\"", list["\"message\""])
}

func TestLoginUser(t *testing.T) {
	NewServer()

	// hashPass, err := hash.HashPassword("ssssssss")
	// if err != nil {
	// 	t.Fatal("failed to password hash: ", err)
	// }

	query := fmt.Sprintf(`
		mutation {
			loginUser(username: %s, password: %s) {
				user_id
				username
				OK
			}
	}`, "\"x\"", "\"x\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"loginUser\"")
	require.Contains(t, string(result), "true")
	require.Equal(t, "\"x\"", list["\"username\""])
}
