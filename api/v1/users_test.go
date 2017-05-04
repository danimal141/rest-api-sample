package v1_test

import (
	"net/http"
	"testing"
)

func TestUsersIndex(t *testing.T) {
	url := server.URL + "/api/v1/users?pagination=false"
	res, err := http.Get(url)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestUsersShow(t *testing.T) {
	url := server.URL + "/api/v1/users/1"
	res, err := http.Get(url)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}
