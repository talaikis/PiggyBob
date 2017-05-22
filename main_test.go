package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHomeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // record response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(TestHomeHandler)

    // e.g. GET /api/projects?page=1&per_page=100
    req, err := http.NewRequest("GET", "/api/projects",
        // Note: url.Values is a map[string][]string
        url.Values{"page": {"1"}, "per_page": {"100"}})
    if err != nil {
        t.Fatal(err)
    }

    // Our handler might also expect an API key.
    req.Header.Set("Authorization", "Bearer abc123")

    // make request
    handler.ServeHTTP(rr, req)

    // check if 200
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    /*expected := `{"alive": true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }*/
}

// more:
// https://elithrar.github.io/article/testing-http-handlers-go/
