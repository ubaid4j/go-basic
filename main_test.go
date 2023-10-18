package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Create a GET request to /albums
	req, _ := http.NewRequest("GET", "/albums", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// Check the response status code
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, resp.Code)
	}

	// Check the response body
	var response []album
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Check the number of albums in the response
	if len(response) != len(albums) {
		t.Errorf("Expected %d albums in response, but got %d", len(albums), len(response))
	}
}

func TestPostAlbums(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.POST("/albums", postAlbums)

	// Create a new album to add
	newAlbum := album{
		ID:     "3",
		Title:  "Red Train",
		Artist: "John",
		Price:  55.5,
	}

	// Convert the new album to JSON
	newAlbumJSON, err := json.Marshal(newAlbum)
	if err != nil {
		t.Errorf("Error converting new album to JSON: %v", err)
	}

	// Create a POST request to /albums with the new album as JSON payload
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(newAlbumJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// Check the response status code
	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status %d, but got %d", http.StatusCreated, resp.Code)
	}

	// Check the response body
	var response album
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Check if the added album matches the new album
	if response != newAlbum {
		t.Errorf("Expected the response to match the new album, but got a different response.")
	}
}

func TestMain(m *testing.M) {
	// You can set up any test-specific configurations or fixtures in this function.
	// For example, if you need to set up a test database, HTTP server, or other resources.
	// After setting up the environment, you can call m.Run() to run the tests.
	m.Run()
}
