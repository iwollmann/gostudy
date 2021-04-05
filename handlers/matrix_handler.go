package handlers

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iwollmann/gostudy/service"
)

type matrixHandler struct {
	service *service.MatrixService
}

func NewMatrixHandler() *matrixHandler {
	return &matrixHandler{
		service.NewMatrixService(),
	}
}

type myKey string

func (mc matrixHandler) AcceptOnlyPostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if (r.Method != http.MethodPost) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mc matrixHandler) ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10<<10) // limit your max input length to 10Mb!
		if err !=nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("csv max length is 10 Mb"))
			return
		}


		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
			return
		}

		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
			return
		}

		matrix, err := mc.validate(records)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
			return
		}
		
		ctx := context.WithValue(r.Context(), myKey("file"), matrix)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mc matrixHandler) Echo(w http.ResponseWriter, r *http.Request){
	matrix, err := matrixFromContext(r.Context())

	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
	} else {
		fmt.Fprint(w, mc.service.Echo(matrix))
	}
}

func (mc matrixHandler) Invert(w http.ResponseWriter,r *http.Request){
	matrix, err := matrixFromContext(r.Context())
	
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
	} else {
		fmt.Fprint(w, mc.service.Invert(matrix))
	}
}
func (mc matrixHandler) Flatten(w http.ResponseWriter, r *http.Request){
	matrix, err := matrixFromContext(r.Context())
	
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
	} else {
		result := mc.service.Flatten(matrix)

		fmt.Fprintf(w, "%s\n", result)
	}
}

func (mc matrixHandler) Sum(w http.ResponseWriter, r *http.Request){
	matrix, err := matrixFromContext(r.Context())

	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
	} else {
		sum := mc.service.Sum(matrix)
		fmt.Fprintf(w, "%s\n", sum)
	}
}

func (mc matrixHandler) Multiply(w http.ResponseWriter, r *http.Request){
	matrix, err := matrixFromContext(r.Context())
	
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
	} else {
		mult := mc.service.Multiply(matrix)

		fmt.Fprintf(w, "%s\n", mult)
	}
}

func (mc matrixHandler) validate(mx [][]string) ([][]int ,error){
	rowsCount := len(mx)
	mx2 := make([][]int, rowsCount);
	if rowsCount == 0 {
		return mx2, errors.New("empty matrix")
	}

	for i := range mx {
		if (rowsCount != len(mx[i])) {
			return mx2, errors.New("this is not a perfect matrix")
		}
		mx2[i] = make([]int, rowsCount)

		for j, v := range mx[i] {
			vInt, err := strconv.ParseInt(v,10,64);
			if err != nil {
				return mx2, fmt.Errorf("invalid value %v at [%v][%v]", v, i, j)
			}

			mx2[i][j] = int(vInt);
		}
	}

	return mx2, nil
}

func matrixFromContext(ctx context.Context) ([][]int, error) {
	v := ctx.Value(myKey("file"))
	if v == nil {
	 return nil,errors.New("unable to retrieve matrix from context")
	}
	return v.([][]int), nil
   }