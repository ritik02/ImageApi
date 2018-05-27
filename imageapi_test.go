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
     expected := "No such file found!"
     actual := strings.Contains(rw.Body.String(),expected)
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v %v", actual, expected, rw.Body.String())
     }
}

func TestWidthandHeightType(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test.jpg&width=abc&height=500a", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "Invalid Parameter Value"
     actual := strings.Contains(rw.Body.String(),expected)
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v %v", actual, expected, rw.Body.String())
     }
}

func TestWidthandHeightRange(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test.jpg&width=400&height=5000", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "Too Big Size Parameters"
     actual := strings.Contains(rw.Body.String(),expected)
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
     }
}

func TestFileResponseAndContentTypeIfValidParameters(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test.jpg&width=500&height=400", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "image/jpeg"
     actual := strings.Contains(rw.Header().Get("Content-Type"),expected)
     if actual == false || len(rw.Body.String()) < 1 {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected,rw.Header().Get("Content-type"))
     }
}

func TestInvalidFileType(t *testing.T) {
     req := httptest.NewRequest(http.MethodGet, "/api/resize?file=text.txt&width=500&height=400", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "Invalid Image Type"
     actual := strings.Contains(rw.Body.String(), expected)
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v v", actual, expected)
     }
}

func TestForPNGImages(t *testing.T) {
  req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test2.png&width=500&height=400", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "image/png"
     actual := strings.Contains(rw.Header().Get("Content-Type"),expected)
     if actual == false || len(rw.Body.String()) < 1 {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected,rw.Header().Get("Content-type"))
     }
}

func TestForCachedFile(t *testing.T) {
  req := httptest.NewRequest(http.MethodGet, "/api/resize?file=test2.png&width=500&height=400", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     expected := "image/png"
     actual := strings.Contains(rw.Header().Get("Content-Type"),expected)
     filename := "images/cached/test2.png_final_500_400.png"
     if actual == false || len(rw.Body.String()) < 1 && DoesFileExist(filename) {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected,rw.Header().Get("Content-type"))
     }
}

