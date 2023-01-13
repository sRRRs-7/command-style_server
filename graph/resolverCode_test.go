package graph

import (
	"fmt"
	"testing"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestAdminCreateCode(t *testing.T) {
	code := utils.RandomString(30)

	query := fmt.Sprintf(`
		mutation {
			createCode(code: %s, img: %s, description: %s, performance: %s, star: %v, tags: %v, access: %d) {
				is_error
				message
			}
	}`, fmt.Sprintf("\"%s\"", code), "\"img\"", "\"description\"", "\"performance\"", []int{1}, []string{"\"go\""}, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/admin/query", token)
	fmt.Println(string(result))

	require.Equal(t, list["\"data\""], "\"createCode\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateCode OK\"", list["\"message\""])
}

func TestCreateCode(t *testing.T) {
	code := utils.RandomString(30)

	query := fmt.Sprintf(`
		mutation {
			createCode(code: %s, img: %s, description: %s, performance: %s, star: %v, tags: %v, access: %d) {
				is_error
				message
			}
	}`, fmt.Sprintf("\"%s\"", code), "\"img\"", "\"description\"", "\"performance\"", []int{6, 8}, []string{"\"go\""}, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}

	token := CreateToken(t)
	_, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(string(result))

	require.Equal(t, list["\"data\""], "\"createCode\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"CreateCode OK\"", list["\"message\""])
}

func TestUpdateCode(t *testing.T) {
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

	token := CreateToken(t)
	arr, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"updateCodes\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateCodes OK\"", list["\"message\""])
}

func TestUpdateAccess(t *testing.T) {
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

	token := CreateToken(t)
	arr, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"updateAccess\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"UpdateAccess OK\"", list["\"message\""])
}

func TestDeleteCode(t *testing.T) {
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

	token := CreateToken(t)
	arr, list, result := NewCookieRequest(t, q, "http://localhost:8080/query", token)
	fmt.Println(arr)

	require.Equal(t, list["\"data\""], "\"deleteCode\"")
	require.Contains(t, string(result), "false")
	require.Equal(t, "\"DeleteCode OK\"", list["\"message\""])
}

func TestGetAllCodes(t *testing.T) {
	id := 5

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
		}`, id)

	q := struct {
		Query string
	}{
		Query: query,
	}

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getCode\"")
	require.Equal(t, fmt.Sprintf("\"%d\"", id), list["\"id\""])
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

func TestGetAllCodesByTagSearch(t *testing.T) {
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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"getAllCodesByTag\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetAllCodesByKeyword(t *testing.T) {
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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesByKeyword\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetAllCodesSortedStar(t *testing.T) {
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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesByKeyword\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetAllCodesSortedAccess(t *testing.T) {
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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "\"GetAllCodesSortedAccess\"")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}

func TestGetAllOwnCodes(t *testing.T) {
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

	_, list, result := NewRequest(t, q, "http://localhost:8080/query", "")
	fmt.Println(string(result))
	fmt.Println(list)

	require.Equal(t, list["\"data\""], "")
	require.Contains(t, string(result), "cookie")
	require.Contains(t, string(result), "[")
	require.Contains(t, string(result), "]")
}
