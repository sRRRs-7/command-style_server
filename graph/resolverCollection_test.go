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
	"github.com/stretchr/testify/require"
)

func TestCreateAdminCollection(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			createAdminCollection(user_id: %d, code_id: %d) {
				is_error
				message
			}
		}`, 2, 6)

	q := struct {
		Query string
	}{
		Query: query,
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	token := CreateAdminToken(t)
	req, _ := http.NewRequest("POST", "http://localhost:8080/admin/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.AdminCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			CreateAdminCollection struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.CreateAdminCollection.Message == "" {
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
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "violates foreign key constraint"))
		require.Equal(t, error.Errors[0].Path[0], "createAdminCollection")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.CreateAdminCollection.IsError, false)
		require.Equal(t, res.Data.CreateAdminCollection.Message, "CreateAdminCollection OK")
	}
}

func TestCreateCollection(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			createCollection(code_id: %d) {
				is_error
				message
			}
		}`, 1)

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
			CreateCollection struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.CreateCollection.Message == "" {
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
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "violates foreign key constraint"))
		require.Equal(t, error.Errors[0].Path[0], "createCollection")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.CreateCollection.IsError, false)
		require.Equal(t, res.Data.CreateCollection.Message, "CreateCollection OK")
	}
}

func TestGetCollection(t *testing.T) {
	r := GinTestRouter()

	id := 7
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
	req, _ := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", resolver.config.RedisCookieKey, token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Println(w.Body)

	var res struct {
		Data struct {
			GetCollection model.CodeWithCollectionID
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.GetCollection.ID == "" {
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
		require.Equal(t, error.Errors[0].Path[0], "getCollection")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, fmt.Sprintf("%d", id), res.Data.GetCollection.ID)
	}
}

func TestDeleteCollection(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			deleteCollection(id: %d) {
				is_error
				message
			}
		}`, 2)

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
			DeleteCollection struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.DeleteCollection.Message == "" {
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
		require.Equal(t, true, strings.Contains(error.Errors[0].Message, "violates foreign key constraint"))
		require.Equal(t, error.Errors[0].Path[0], "deleteCollection")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.DeleteCollection.IsError, false)
		require.Equal(t, res.Data.DeleteCollection.Message, "DeleteCollection OK")
	}
}

func TestGetAllCollection(t *testing.T) {
	r := GinTestRouter()

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
			GetAllCollection []model.CodeWithCollectionID
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCollection) >= 0)
}

func TestGetAllCollectionBySearch(t *testing.T) {
	r := GinTestRouter()

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
		}`, "\"go\"", 10, 0)

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
			GetAllCollectionBySearch []model.CodeWithCollectionID
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllCollectionBySearch) >= 0)
}
