package main

import "fmt"

// https://mayurwadekar2.medium.com/golangs-go-build-command-f471a5e8535d
func main() {
	fmt.Println("Hi, World!")
}

/*
go build -n sourcefile.go
It helps to understand what commands will be execute while building a binary. (This only shows commands and not execute it).
*/
