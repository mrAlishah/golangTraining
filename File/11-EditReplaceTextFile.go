package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://stackoverflow.com/a/41369186
// https://stackoverflow.com/a/70781074
// https://www.golangdev.in/2021/08/read-file-using-fseek-or-seek-in-golang.html
func main() {
	WriteFileByWriteSlice("03.txt")
	EditFileLineByLine("03.txt", "2", "9")
	ReadFileLineByLine("03.txt")
}

func EditFileLineByLine(filename string, oldstr string, newstr string) {
	fmt.Println("\nEditFileLineByLine ------------------------------")
	src, err := os.Open(filename)
	//]src, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	var lenSeek int64 = 0
	scanner := bufio.NewScanner(src)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)

		if line == oldstr {
			dst.Seek(lenSeek, 1)
			dst.WriteString(newstr)
		}

		lenSeek += int64(len(line) + 1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func WriteFileByWriteSlice(filename string) {
	fmt.Println("\nWriteFileByWriteSlice ------------------------------")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := []string{"1", "2", "3", "4", "5", "6"}
	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
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
