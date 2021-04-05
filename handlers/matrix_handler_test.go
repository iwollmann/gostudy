package handlers

import (
	"net/http"
	"net/http/httptest"
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

// func TestValidateMiddleware(t *testing.T) {
// 	postData :=
// 	`	--xxx
// 		Content-Disposition: form-data; name="file"

// 		value1
// 		--xxx
// 		Content-Disposition: form-data; name="wrong"

// 		value2
// 		--xxx
// 		Content-Disposition: form-data; name="file"; filename="file"
// 		Content-Type: application/octet-stream
// 		Content-Transfer-Encoding: binary

// 		binary data
// 		--xxx--
// 		`

// 	req := &Request{
// 		Method: "POST",
// 		Header: Header{"Content-Type": {`multipart/form-data; boundary=xxx`}},
// 		Body:   io.NopCloser(strings.NewReader(postData)),
// 	}
// 	tests := []struct {
// 		name string
// 		file string
// 		expected int
// 	}{
// 		{ "empty file", "", 404 },
// 		{ "valid file", "1,2,3\n4,5,6\n,7,8,9", 200 },
// 		{ "invalid file", "matrix.txt", 404 },
// 	}

// 	mc := NewMatrixHandler()

// 	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	})

// 	handler := mc.ValidateMiddleware(testHandler)

// 	io.NopCloser()
	
// 	for _, tt := range tests {
// 		test := strings.NewReader(tt.file)

// 		req, err := http.NewRequest("POST", "/", test)
// 		// req.Header.Add("Content-Type", mwriter.FormDataContentType())
// 		if err != nil {
// 			t.Fatal(err)
// 		}
	
// 		rr := httptest.NewRecorder()
// 		handler.ServeHTTP(rr, req)

//         got :=rr.Result().StatusCode
//         if got != tt.expected {
// 			t.Errorf("case %v returned unexpected status code: got %v expected %v",
//             tt.name, got, tt.expected)
//         }
//     }
// }

// func TestFlatten(t *testing.T) {
// 	flatTests := []struct {
// 		matrix [][]int
// 		expected string
// 	}{
// 		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "1,2,3,4,5,6,7,8,9\n" },
// 		{ [][]int{ {1,2}, {3,4} }, "1,2,3,4\n" },
// 	}

// 	req, err := http.NewRequest("POST", "/flatten", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

// 	mc := NewMatrixHandler()

// 	for _, tt := range flatTests {
// 		ctx := req.Context()
// 		ctx = context.WithValue(ctx, myKey("file"), tt.matrix)
		
// 		req = req.WithContext(ctx)
// 		rr := httptest.NewRecorder()
// 		mc.Flatten(rr,req)

//         got := rr.Body.String()
//         if got != tt.expected {
// 			t.Errorf("handler returned unexpected body: got %v want %v",
//             got, tt.expected)
//         }
//     }
// }
