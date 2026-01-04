package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CHILLERAN/Snippetbox/internal/assert"
)

func TestCommonHeaders(t *testing.T) {
	//Arrange
	responseRecorder := httptest.NewRecorder()	
	expectedContentSecurityPolicy := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	expectedReferrerPolicy := "origin-when-cross-origin"
	expectedContentTypeOptions := "nosniff"
	expectedXFrameOptions := "deny"
	expectedXSSProtection := "0"
	expectedServer := "Go"
	
	//Act
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

	commonHeaders(next).ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
    defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
    if err != nil {
        t.Fatal(err)
    }
	
    body = bytes.TrimSpace(body)

	//Assert
	assert.Equal(t, result.Header.Get("Content-Security-Policy"), expectedContentSecurityPolicy)	
	assert.Equal(t, result.Header.Get("Referrer-Policy"), expectedReferrerPolicy)
	assert.Equal(t, result.Header.Get("X-Content-Type-Options"), expectedContentTypeOptions)
	assert.Equal(t, result.Header.Get("X-Frame-Options"), expectedXFrameOptions)
	assert.Equal(t, result.Header.Get("X-XSS-Protection"), expectedXSSProtection)
	assert.Equal(t, result.Header.Get("Server"), expectedServer)
	assert.Equal(t, result.StatusCode, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}