package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "strings"
)

func init() {
     http.HandleFunc("/api/resize", ImageHandler)
}

func TestCorrectResponse(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     if rw.Code != 200 {			
     t.Fatalf("Expected 200 response code, but got: %v\n", rw.Code)
}
}

func TestParamters(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "Invalid Paramters"
     actual := strings.Contains(rw.Body.String(),"Invalid Parameters")
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestFileAbsence(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test3.jpg&width=600&height=500", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "No Such file found!"
     actual := strings.Contains(rw.Body.String(),expected)
     if actual == true {
     	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
     }

}
