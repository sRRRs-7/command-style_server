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

	"github.com/stretchr/testify/require"
)

func TestCreateAdminUser(t *testing.T) {
	r := GinTestRouter()

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

	fmt.Println(w.Body)

	var res struct {
		Data struct {
			CreateAdminUser struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.CreateAdminUser.Message == "" {
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
		require.Equal(t, error.Errors[0].Path[0], "createAdminUser")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.CreateAdminUser.IsError, false)
		require.Equal(t, res.Data.CreateAdminUser.Message, "CreateAdminUser OK")
	}

}

func TestGetAdminUser(t *testing.T) {
	r := GinTestRouter()

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
			GetAdminUser struct {
				Is_Username bool
				Is_Password bool
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.GetAdminUser.Is_Username && res.Data.GetAdminUser.Is_Password {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, true, res.Data.GetAdminUser.Is_Username)
		require.Equal(t, true, res.Data.GetAdminUser.Is_Password)

	} else {
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
		require.Equal(t, error.Errors[0].Path[0], "getAdminUser")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	}

}
