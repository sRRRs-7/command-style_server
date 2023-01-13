package graph

import (
	"fmt"
	"testing"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUserResolver(t *testing.T) {
	NewServer()

	username := utils.RandomString(10)
	email := utils.RandomEmail()

	query := fmt.Sprintf(`
		mutation {
			createUser(username: %s, password: %s, email: %s, sex: %s, date_of_birth: %s) {
				is_error
				message
			}
	}`, fmt.Sprintf("\"%s\"", username), "\"srrrs\"", fmt.Sprintf("\"%s\"", email), "\"man\"", "\"1996-13-45\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://127.0.0.1:8080/query", "")

	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"createUser\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateUser OK\"", list["\"message\""])
}

func TestLoginUser(t *testing.T) {
	NewServer()

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

	arr, list, result := NewRequest(t, q, "http://127.0.0.1:8080/query", "")

	fmt.Println(arr)
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"loginUser\"")
	require.Contains(t, string(result), "true")
	require.Equal(t, "\"x\"", list["\"username\""])
}
