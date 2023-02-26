package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const fileName = "01.txt"

func main() {
	ReadFile(fileName)
	ReadFileByBuffer(fileName)
	ReadFileByStringsBuilder(fileName)
	ReadFileByEmbedPkg()
}

func ReadFile(filename string) {
	fmt.Println("\nread file by os.ReadFile ------------------------------")
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//defer f.Close() //don't need

	//contents := f.String() //cast is different
	contents := string(f) //[]byte

	fmt.Print(contents)
}

func ReadFileByBuffer(filename string) {
	fmt.Println("\nread file by bytes.Buffer------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	contents := buf.String() //*bytes.Buffer

	fmt.Print(contents)
}

func ReadFileByStringsBuilder(filename string) {
	fmt.Println("\nread file by strings.Builder------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sb := new(strings.Builder)
	io.Copy(sb, f)
	contents := sb.String() //*strings.Builder

	fmt.Print(contents)
}

//go:embed 01.txt
var contents string

// embed cannot apply to var inside func
func ReadFileByEmbedPkg() {
	fmt.Println("\nread file by embed PKG------------------------------")

	fmt.Print(contents) // print the content as a 'string'
}
