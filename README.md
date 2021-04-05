[![Build Status](https://github.com/iwollmann/gostudy/workflows/go/badge.svg)](https://github.com/iwollmann/gostudy/actions?workflow=go)

# Golang WebService Sample
Simple webservice written in Golang without any 3rd party library to process a matrix from csv file and perform some operations.

### Operations:

1. Echo
    - Return the matrix as a string in matrix format.

2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted

3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
 
4. Sum
    - Return the sum of the integers in the matrix

5. Multiply
    - Return the product of the integers in the matrix


### Examples
Test API with curl
```
curl -F 'file=@matrix_sample.csv' -v http://localhost:8080/echo

1,2,3
4,5,6
7,8,9
```

```
curl -F 'file=@matrix_sample.csv' -v http://localhost:8080/flatten

1,2,3,4,5,6,7,8,9
```

```
curl -F 'file=@matrix_sample.csv' -v http://localhost:8080/invert

1,4,7
2,5,8
3,6,9
```

```
curl -F 'file=@matrix_sample.csv' -v http://localhost:8080/sum

45
```

```
curl -F 'file=@matrix_sample.csv' -v http://localhost:8080/multiply

362880
```


## Installation
```
  go get github.com/iwollmann/gostudy
```

## Run in Docker
- Docker Compose
```
docker-compose up
```

## Makefile
- Build binary to ./build/
```
make build
```
- Run tests
```
make test
```
- Clean up tests and binary files
```
make clean
```
- Build For Linux
```
make build-linux
```

## Go version
``go1.16.2 linux/amd64``

