package service

import (
	"fmt"
	"strconv"
	"strings"
)

type MatrixService struct {}

func NewMatrixService() *MatrixService {
	return &MatrixService{}
}

func(mx MatrixService) Echo(matrix [][]int) string {
	return format(matrix)
}

func (mx MatrixService) Invert(matrix [][]int) string {
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return format(matrix)
}
 
func (mx MatrixService) Flatten(matrix [][]int) string {
	var response string
	for _, row := range matrix {
		var sRow []string
		for _, i := range row {
			sRow = append(sRow, strconv.Itoa(i))
		}
		if (response == "") {
			response = strings.Join(sRow, ",")
		} else {
			response = fmt.Sprintf("%s,%s", response, strings.Join(sRow, ","))
		}
	}

	return response
}

func (mx MatrixService) Sum(matrix [][]int) string {
	sum := 0
	for _, row := range matrix {
		for _, v := range row {
			sum += v
		}
	}

	return strconv.Itoa(sum)
}

func (mx MatrixService) Multiply(matrix [][]int) string {
	mult := 1
	for _, row := range matrix {
		for _, v := range row {
			fmt.Printf("%v", v)
			mult *= v
		}
	}

	return strconv.Itoa(mult)
}

func format(matrix [][]int) string {
	var response string
	for _, row := range matrix {
		var sRow []string
		for _, i := range row {
			sRow = append(sRow, strconv.Itoa(i))
		}
		response = fmt.Sprintf("%s%s\n", response, strings.Join(sRow, ","))
	}
	return response
}