package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAdminUser(t *testing.T) {
	NewServer()

	// hashPass, err := hash.HashPassword("ssssssss")
	// if err != nil {
	// 	t.Fatal("failed to password hash: ", err)
	// }

	query := fmt.Sprintf(`
		mutation {
			createAdminUser(username: %s, password: %s) {
				is_error
				message
			}
	}`, "\"ssssss\"", "\"ssssssss\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createAdminUser\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateMedia OK\"", list["\"message\""])
}

// fix hash password logic
func TestGetAdminUser(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
	mutation {
		getAdminUser(username: %s, password: %s) {
			is_username
			is_password
		}
	}`, "\"xvlbz\"", "\"gbaicmra\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAdminUser\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateMedia OK\"", list["\"message\""])
}
