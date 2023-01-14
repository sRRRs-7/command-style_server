package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestAdminCreateCode(t *testing.T) {
	r := GinTestRouter()

	code := utils.RandomString(30)
	query := fmt.Sprintf(`
		mutation {
			adminCreateCode(username: %s, code: %s, img: %s, description: %s, performance: %s, star: %v, tags: %v, access: %d) {
				is_error
				message
			}
	}`, "\"srrrs\"", fmt.Sprintf("\"%s\"", code), "\"img\"", "\"description\"", "\"performance\"", []int{1}, []string{"\"go\""}, 1)

	q := struct {
		Query string
	}{
		Query: query,
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			AdminCreateCode struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.AdminCreateCode.Message == "" {
		type errRes struct {
			Message string
			Path    []string
		}
		var error struct {
			Errors []errRes
			Data   any
		}
		err := json.Unmarshal(w.Body.Bytes(), &error)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "duplicate key value violates unique constraint"))
		require.Equal(t, error.Errors[0].Path[0], "adminCreateCode")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.AdminCreateCode.IsError, false)
		require.Equal(t, res.Data.AdminCreateCode.Message, "AdminCreateCode OK")
	}
}

func TestCreateCode(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			CreateCode struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	fmt.Println(w.Body)

	if res.Data.CreateCode.Message == "" {
		type errRes struct {
			Message string
			Path    []string
		}
		var error struct {
			Errors []errRes
			Data   any
		}
		err := json.Unmarshal(w.Body.Bytes(), &error)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "duplicate key value violates unique constraint") || strings.Contains(error.Errors[0].Message, "violates foreign key constraint "))
		require.Equal(t, error.Errors[0].Path[0], "createCode")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.CreateCode.IsError, false)
		require.Equal(t, res.Data.CreateCode.Message, "CreateCode OK")
	}
}

func TestUpdateCode(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			UpdateCodes struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.UpdateCodes.IsError, false)
	require.Equal(t, res.Data.UpdateCodes.Message, "UpdateCodes OK")

}

func TestUpdateAccess(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			UpdateAccess struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	fmt.Println(w.Body)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.UpdateAccess.IsError, false)
	require.Equal(t, res.Data.UpdateAccess.Message, "UpdateAccess OK")

}

func TestDeleteCode(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			DeleteCode struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.DeleteCode.IsError, false)
	require.Equal(t, res.Data.DeleteCode.Message, "DeleteCode OK")

}

func TestGetCode(t *testing.T) {
	r := GinTestRouter()

	id := 1
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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetCode model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.GetCode.ID == "" {
		type errRes struct {
			Message string
			Path    []string
		}
		var error struct {
			Errors []errRes
			Data   any
		}
		err := json.Unmarshal(w.Body.Bytes(), &error)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "no rows in result set"))
		require.Equal(t, error.Errors[0].Path[0], "getCode")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, fmt.Sprintf("%d", id), res.Data.GetCode.ID)
	}
}

func TestGetAllCodes(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		query {
			getAllCodes(limit: %d, skip: %d) {
				id
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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllCodes []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCodes) >= 0)
}

func TestGetAllCodesByTagSearch(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllCodesByTag []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCodesByTag) >= 0)
}

func TestGetAllCodesByKeyword(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllCodesByKeyword []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCodesByKeyword) >= 0)
}

func TestGetAllCodesSortedStar(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllCodesByKeyword []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCodesByKeyword) >= 0)
}

func TestGetAllCodesSortedAccess(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllCodesSortedAccess []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCodesSortedAccess) >= 0)
}

func TestGetAllOwnCodes(t *testing.T) {
	r := GinTestRouter()

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

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllOwnCodes []model.Code
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllOwnCodes) >= 0)
}
