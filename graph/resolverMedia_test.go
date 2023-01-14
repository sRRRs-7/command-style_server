package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/stretchr/testify/require"
)

func TestCreateMedia(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			createMedia(title: %s, contents: %s, img: %s) {
				is_error
				message
			}
	}`, "\"title\"", "\"contents\"", "\"img\"")

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
			CreateMedia struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.CreateMedia.IsError, false)
	require.Equal(t, res.Data.CreateMedia.Message, "CreateMedia OK")
}

func TestGetAllMedia(t *testing.T) {
	query := fmt.Sprintf(`
		query {
			getAllMedia(limit: %d, skip: %d) {
				id
				title
				contents
				img
				created_at
				updated_at
			}
	}`, 10, 0)

	r := GinTestRouter()

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
			GetAllMedia []model.Media
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, true, len(res.Data.GetAllMedia) >= 0)
}

func TestGetMediaResolver(t *testing.T) {
	r := GinTestRouter()

	id := 1
	query := fmt.Sprintf(`
		mutation {
			getMedia(id: %d) {
				id
				title
				contents
				img
				created_at
				updated_at
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
	req, _ := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(body.Bytes()))
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res struct {
		Data struct {
			GetAllMedia model.Media
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	if res.Data.GetAllMedia.ID == "" {
		type errRes struct {
			Message string
			Path    []string
		}
		var error struct {
			Errors []errRes
			Data   any
		}
		err := json.Unmarshal(w.Body.Bytes(), &error)
		fmt.Println(error)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, error.Errors[0].Message, "GetMedia error: failed to get a media: no rows in result set")
		require.Equal(t, error.Errors[0].Path[0], "getMedia")
		require.Equal(t, reflect.TypeOf(error.Data), nil)
	} else {
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, res.Data.GetAllMedia.ID, id)
		require.Equal(t, reflect.TypeOf(res.Data.GetAllMedia.Title), reflect.String)
	}
}

func TestUpdateMedia(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			updateMedia(id: %s, title: %s, contents: %s, img: %s) {
				is_error
				message
			}
	}`, "\"1\"", "\"title\"", "\"contents\"", "\"img\"")

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
			UpdateMedia struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.UpdateMedia.IsError, false)
	require.Equal(t, res.Data.UpdateMedia.Message, "UpdateMedia OK")
}

func TestDeleteMedia(t *testing.T) {
	r := GinTestRouter()

	query := fmt.Sprintf(`
		mutation {
			deleteMedia(id: %d) {
				is_error
				message
			}
	}`, 3)

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
			DeleteMedia struct {
				IsError bool
				Message string
			}
		}
	}
	err := json.Unmarshal(w.Body.Bytes(), &res)

	fmt.Println(w.Body)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, res.Data.DeleteMedia.IsError, false)
	require.Equal(t, res.Data.DeleteMedia.Message, "DeleteMedia OK")
}
