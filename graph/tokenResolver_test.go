package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	CreateToken(t)
}

func CreateToken(t *testing.T) string {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			createToken(username: %s)
	}`, "\"srrrs\"")

	q := struct {
		Query string
	}{
		Query: query,
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			CreateToken string
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)

	return res.Data.CreateToken
}

func TestAdminCreateToken(t *testing.T) {
	CreateAdminToken(t)
}

func CreateAdminToken(t *testing.T) string {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			createAdminToken(username: %s, password: %s)
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
			CreateAdminToken string
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)

	return res.Data.CreateAdminToken
}
