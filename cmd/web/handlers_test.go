package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CHILLERAN/Snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	//Unit Test
	/*responseRecorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

	ping(responseRecorder, req)

	result := responseRecorder.Result()
    defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
    if err != nil {
        t.Fatal(err)
    }
	
    body = bytes.TrimSpace(body)

	assert.Equal(t, result.StatusCode, http.StatusOK)
    assert.Equal(t, string(body), "OK")*/

	//End to End test
	app := &Application{
        logger: slog.New(slog.DiscardHandler),
    }

	ts := httptest.NewTLSServer(app.routes())
    defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/ping", nil)
    if err != nil {
        t.Fatal(err)
    }

	res, err := ts.Client().Do(req)
    if err != nil {
        t.Fatal(err)
    }

    defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
    if err != nil {
        t.Fatal(err)
    }

    body = bytes.TrimSpace(body)

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}