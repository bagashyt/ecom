package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseJSON(t *testing.T) {
	type Payload struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name        string
		requestBody []byte
		wantPayload Payload
		wantErr     bool
	}{
		{
			name:        "Valid JSON",
			requestBody: []byte(`{"name": "Alice", "age": 30}`),
			wantPayload: Payload{Name: "Alice", Age: 30},
			wantErr:     false,
		},
		{
			name:        "Invalid JSON",
			requestBody: []byte(`{"name": "Alice", "age": "thirty"}`),
			wantPayload: Payload{},
			wantErr:     true,
		},
		{
			name:        "Empty body",
			requestBody: nil,
			wantPayload: Payload{},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", bytes.NewReader(tt.requestBody))
			var gotPayload Payload

			err := ParseJSON(req, &gotPayload)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && gotPayload != tt.wantPayload {
				t.Errorf("ParseJSON() got = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}

func TestWriteJSON(t *testing.T) {
	type Response struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	tests := []struct {
		name           string
		status         int
		responseBody   Response
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid JSON response",
			status:         http.StatusOK,
			responseBody:   Response{Message: "Success", Code: 200},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Success","code":200}`,
		},
		{
			name:           "Error JSON response",
			status:         http.StatusBadRequest,
			responseBody:   Response{Message: "Bad Request", Code: 400},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Bad Request","code":400}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			err := WriteJSON(rr, tt.status, tt.responseBody)
			if err != nil {
				t.Fatalf("WriteJSON() error = %v", err)
			}

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("WriteJSON() status code = %v, want %v", status, tt.expectedStatus)
			}

			// Check the Content-Type header
			if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
				t.Errorf("WriteJSON() content type = %v, want application/json", contentType)
			}

			// Check the response body
			var responseBody map[string]any
			if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("Response body is not valid JSON: %v", err)
			}

			var expectedBody map[string]any
			if err := json.Unmarshal([]byte(tt.expectedBody), &expectedBody); err != nil {
				t.Fatalf("Expected body is not valid JSON: %v", err)
			}

			// if !equal(responseBody, expectedBody) {
			// 	t.Errorf("WriteJSON() body = %v, want %v", responseBody, expectedBody)
			// }
		})
	}
}
