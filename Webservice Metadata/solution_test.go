package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetCustomerMetadata(t *testing.T) {
	handler := NewHandler()

	t.Run("valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customer?id=0", nil)
		w := httptest.NewRecorder()

		handler.HandleGetCustomerMetadata(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200, got %d", resp.StatusCode)
		}

		var result Response
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if result.Id != 0 || result.Name != "Alex" {
			t.Errorf("Unexpected response: %+v", result)
		}
	})

	t.Run("missing ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customer", nil)
		w := httptest.NewRecorder()

		handler.HandleGetCustomerMetadata(w, req)

		if w.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Result().StatusCode)
		}
	})

	t.Run("non-numeric ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customer?id=abc", nil)
		w := httptest.NewRecorder()

		handler.HandleGetCustomerMetadata(w, req)

		if w.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Result().StatusCode)
		}
	})

	t.Run("nonexistent ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customer?id=999", nil)
		w := httptest.NewRecorder()

		handler.HandleGetCustomerMetadata(w, req)

		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Result().StatusCode)
		}
	})
}

func TestHandleGetAllCustomerMetadata(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodGet, "/customers", nil)
	w := httptest.NewRecorder()

	handler.HandleGetAllCustomerMetadata(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	var list []Response
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	expectedLength := len(handler.custDb)
	if len(list) != expectedLength {
		t.Errorf("Expected %d entries, got %d", expectedLength, len(list))
	}

	for _, entry := range list {
		if entry.Name == "" {
			t.Errorf("Empty name in entry: %+v", entry)
		}
		if _, ok := handler.custDb[entry.Id]; !ok {
			t.Errorf("Unexpected ID in list: %d", entry.Id)
		}
	}
}
