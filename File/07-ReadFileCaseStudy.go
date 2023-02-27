package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// https://zetcode.com/golang/readfile/?utm_content=cmp-true
func main() {
	ReadFileIntoString("01.txt")
	ReadFileLineByLine("01.txt")
	ReadFileByWords("01.txt")
	ReadFileInChunks("01.txt")
	ReadBinaryFile("02.bin")
}

func ReadFileIntoString(filename string) {
	fmt.Println("\nReadFileIntoString ------------------------------")
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func ReadFileLineByLine(filename string) {
	fmt.Println("\nReadFileLineByLine ------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadFileByWords(filename string) {
	fmt.Println("\nReadFileByWords ------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func ReadFileInChunks(filename string) {
	fmt.Println("\nReadFileInChunks ------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		fmt.Print(string(buf[:n]))
	}

	fmt.Println()
}

func ReadBinaryFile(filename string) {
	fmt.Println("\nReadBinaryFile ------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)
	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//we read binary file and print it in hexadecimal format.
		fmt.Printf("%s", hex.Dump(buf))
	}
}
