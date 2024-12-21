package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/StepanShel/YaCalc/application"
	"github.com/StepanShel/YaCalc/pkg/calculation"
)

type RequestBody struct {
	Expression string `json:"expression"`
}

func TestCalcHandSuccess(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()
	requestBody := RequestBody{
		Expression: "1+1",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}
	var response application.ResponseRes
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	expectedResult := 2.000000
	if response.Result != expectedResult {
		t.Fatal("Expected result", expectedResult, "got ", response.Result, response.Result)
	}
}

func TestCalcHandInvalidExpression(t *testing.T) {

	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()
	requestBody := RequestBody{
		Expression: "1+/",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}
	var response application.ResponseError
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != "invalid expression" {
		t.Fatalf("Expected error %v, got %v", "invalid expression", response.Error)
	}
}
func TestCalcHandDivisionByZero(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()
	requestBody := RequestBody{
		Expression: "1/0",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}
	var response application.ResponseError
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calculation.DivByZero.Error() {
		t.Fatalf("Expected error %v, got %v", calculation.DivByZero, response.Error)
	}
}
