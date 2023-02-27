package main

import (
	"fmt"
	"log"
	"os"
)

// https://zetcode.com/golang/writefile/
func main() {
	WriteFileByWriteString("02.txt")
	WriteFileByWriteFile("02.txt")
	WriteFileByWriteAt("02.txt")
	WriteFileByWriteSlice("02.txt")
	WriteFileByFprintln("02.txt")
}
func WriteFileByWriteString(filename string) {
	fmt.Println("\nWriteFileByWriteString ------------------------------")
	//The os.Create creates or truncates the named file. If the file already exists, it is truncated.
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString("New line data\n")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func WriteFileByWriteFile(filename string) {
	fmt.Println("\nWriteFileByWriteFile ------------------------------")
	val := "New line data\n"
	data := []byte(val)

	err := os.WriteFile(filename, data, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}

// First, we write "New line data" to the file with File.Write,
// then we write the " and add another data" next to it with File.WriteAt.
func WriteFileByWriteAt(filename string) {
	fmt.Println("\nWriteFileByWriteAt ------------------------------")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//transform the string to bytes.
	val := "New line data"
	data := []byte(val)

	//write the bytes with Write.
	_, err2 := f.Write(data)
	if err2 != nil {
		log.Fatal(err2)
	}

	val2 := " and add another data\n"
	data2 := []byte(val2)

	//We calculate the length of the previously written string.
	var idx int64 = int64(len(data))

	//write the string and the specified index with WriteAt.
	_, err3 := f.WriteAt(data2, idx)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("done")
}

func WriteFileByWriteSlice(filename string) {
	fmt.Println("\nWriteFileByWriteSlice ------------------------------")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []string{"New", "Line", "Data", "By", "Slice"}
	for _, word := range words {
		_, err := f.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func WriteFileByFprintln(filename string) {
	fmt.Println("\nWriteFileByFprintln ------------------------------")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	const name, age = "Johne Doe", 34

	n, err := fmt.Fprintln(f, name, "is", age, "years old.")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "bytes written")
	fmt.Println("done")
}
