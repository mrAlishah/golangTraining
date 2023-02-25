package main

import (
	"embed"
	"fmt"
)

//https://zetcode.com/golang/embed/
//https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/
func main() {
	ReadFileSingleFile()
	ReadFileMultiFiles()
}

//Go embed a text file
var (
    //go:embed 01.txt
    data string
)
func ReadFileSingleFile() {
	fmt.Println("\nread file by embed Single File ------------------------------")	
 
	fmt.Print(data) // print the content as a 'string'
}		

//Go embed multiple files
var (
	//go:embed *.txt
    f embed.FS
)
func ReadFileMultiFiles() {
	fmt.Println("\nread file by embed Multi Files ------------------------------")	
 
	fmt.Println("\nread file 01.txt =>")
    txt1, _ := f.ReadFile("01.txt")
    fmt.Println(string(txt1))

	fmt.Println("\nread file 03.txt =>")
    txt3, _ := f.ReadFile("03.txt")
    fmt.Println(string(txt3))	

}	