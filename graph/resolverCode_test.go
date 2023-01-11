package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// fix cookie logic
func TestAdminCreateCode(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createCode(code: %s, img: %s, description: %s, performance: %s, star: %v, tags: %v, access: %d) {
				is_error
				message
			}
	}`, "\"ifconfig\"", "\"img\"", "\"description\"", "\"performance\"", []int{1}, []string{"\"go\""}, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, _ := NewRequest(t, q, "http://localhost:8080/admin/query")
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"createCode\"")
	require.Equal(t, true, len(list["\"is_error\""]))
	require.Equal(t, true, len(list["\"message\""]))
}

// fix cookie logic
func TestCreateCode(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			createCode(code: %s, img: %s, description: %s, performance: %s, star: %v, tags: %v, access: %d) {
				is_error
				message
			}
	}`, "\"ifconfig\"", "\"img\"", "\"description\"", "\"performance\"", []int{1, 2}, []string{"\"go\""}, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}
	arr, list, _ := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"createCode\"")
	require.Equal(t, false, len(list["\"is_error\""]))
	require.Equal(t, true, len(list["\"message\""]))
}

func TestUpdateCode(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			updateCodes(id: %d, code: %s, img: %s, description: %s, performance: %s, tags: %v) {
				is_error
				message
			}
		}`, 1, "\"commend\"", "\"img\"", "\"description\"", "\"performance\"", []string{"\"go\""})

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"updateCodes\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateCodes OK\"", list["\"message\""])
}

func TestUpdateAccess(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			updateAccess(id: %d, access: %d) {
				is_error
				message
			}
		}`, 3, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"updateAccess\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateAccess OK\"", list["\"message\""])
}

func TestDeleteCode(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		mutation {
			deleteCode(id: %d) {
				is_error
				message
			}
		}`, 4)

	q := struct {
		Query string
	}{
		Query: query,
	}

	arr, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"deleteCode\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteCode OK\"", list["\"message\""])
}

func TestGetAllCodes(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			getCode(id: %d) {
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
		}`, 4)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getCode\"")
	require.Equal(t, "\"4\"", list["\"id\""])
	require.True(t, 3 <= len(list["\"username\""]))
	require.True(t, 3 <= len(list["\"code\""]))
	require.True(t, 3 <= len(list["\"img\""]))
	require.True(t, 3 <= len(list["\"description\""]))
	require.Contains(t, string(result), "star")
	require.Contains(t, string(result), "go") // tags
	require.Contains(t, string(result), "created_at")
	require.Contains(t, string(result), "updated_at")
	require.Contains(t, string(result), "access")
	require.Contains(t, string(result), "user_id")
}

// fix cookie logic
func TestGetAllCodesByTagSearch(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			getAllCodesByTag(tags: %v, sortBy: %s, limit: %d, skip: %d) {
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
		}`, []string{"\"go\""}, EnumSort.Asc, 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllCodesByTag\"")
	require.True(t, len(string(result)) >= 1000)
}

// fix cookie logic
func TestGetAllCodesByKeyword(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			GetAllCodesByKeyword(keyword: %s, limit: %d, skip: %d) {
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
		}`, "\"gogogogo\"", 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesByKeyword\"")
	require.True(t, len(string(result)) <= 100)
}

// fix cookie logic
func TestGetAllCodesSortedStar(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			GetAllCodesByKeyword(keyword: %s, limit: %d, skip: %d) {
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
		}`, "\"gogogogoggo\"", 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesByKeyword\"")
	require.True(t, len(string(result)) <= 100)
}

// fix cookie logic
func TestGetAllCodesSortedAccess(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			GetAllCodesSortedAccess(limit: %d, skip: %d) {
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
		}`, 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesSortedAccess\"")
	require.True(t, len(string(result)) >= 1000)
}

func TestGetAllOwnCodes(t *testing.T) {
	NewServer()

	query := fmt.Sprintf(`
		query {
			getAllOwnCodes(limit: %d, skip: %d) {
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
		}`, 10, 0)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "")
	require.Contains(t, string(result), "cookie")
}
