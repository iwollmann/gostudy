package handlers

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAcceptsOnlyPOSTMiddleware(t *testing.T) {
	tests := []struct {
		method string
		expected int
	}{
		{ "GET", 404 },
		{ "POST", 200 },
		{ "PUT", 404 },
	}

	mc := NewMatrixHandler()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := mc.AcceptOnlyPostMiddleware(testHandler)
	
	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
	
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

        got :=rr.Result().StatusCode
        if got != tt.expected {
			t.Errorf("%v call returned unexpected status code: got %v expected %v",
            tt.method, got, tt.expected)
        }
    }
}

func TestValidateMiddleware(t *testing.T) {
	tests := []struct {
		name string
		file string
		content string
		expected int
	}{
		{ "wrong file name", "othername", "", 400 },
		{ "empty file", "file", "", 400 },
		{ "valid file", "file", "1,2,3\n4,5,6\n7,8,9", 200 },
		{ "invalid matrix content", "file","1.2,3\n2.3, 4", 400 },
		{ "invalid matrix", "file","1,3,5\n2, 4", 400 },
	}

	mc := NewMatrixHandler()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := mc.ValidateMiddleware(testHandler)
	
	for _, tt := range tests {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormFile(tt.file, "test.csv")
		reader := strings.NewReader(tt.content)
		if err != nil {
			t.Fatal(err)
		}

		_, err = io.Copy(fw, reader)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()
		req, err := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		if err != nil {
			t.Fatal(err)
		}
	
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

        got :=rr.Result().StatusCode
        if got != tt.expected {
			t.Errorf("case %v returned unexpected status code: got %v expected %v",
            tt.name, got, tt.expected)
        }
    }
}

func TestHandlers(t* testing.T) {
	mc := NewMatrixHandler()

	tests := []struct {
		name string
		withContext bool
		handler func(w http.ResponseWriter, r *http.Request)
		expected int
	}{
		{ "Echo", true, mc.Echo, 200 },
		{ "Echo", false, mc.Echo, 500 },
		{ "Invert", true, mc.Invert, 200 },
		{ "Invert", false, mc.Invert, 500 },
		{ "Flatten", true, mc.Flatten, 200 },
		{ "Flatten", false, mc.Flatten, 500 },
		{ "Sum", true, mc.Sum, 200 },
		{ "Sum", false, mc.Sum, 500 },
		{ "Multiply", true, mc.Multiply, 200 },
		{ "Multiply", false, mc.Multiply, 500 },
	}

	
	for _, tt := range tests {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tt.handler(w, r)
			w.WriteHeader(http.StatusOK)
		})

		if (tt.withContext) {
			ctx := req.Context()
    		ctx = context.WithValue(ctx, myKey("file"), [][]int {{}})
			req = req.WithContext(ctx)
		}

		rr := httptest.NewRecorder()
		testHandler.ServeHTTP(rr, req)

        got :=rr.Result().StatusCode
        if got != tt.expected {
			t.Errorf("%v call returned unexpected status code: got %v expected %v",
            tt.name, got, tt.expected)
        }
    }
}