package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

//https://www.jonathan-petitcolas.com/2014/09/25/parsing-binary-files-in-go.html

// this type represnts a record with three fields
type payload struct {
	One   float32
	Two   float64
	Three uint32
}

func main() {
	writeFile()
	readFile()
	//readFileRefactored()
}

// -------------- WriteFile ---------------------------
func writeFile() {
	fmt.Println("\nStart to write binary file ------------------------------")
	file, err := os.Create("02.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {

		s := &payload{
			r.Float32(),
			r.Float64(),
			r.Uint32(),
		}

		var buf bytes.Buffer
		err = binary.Write(&buf, binary.BigEndian, s) //binary.LittleEndian
		//err = binary.Write(file, binary.BigEndian, s) //write to file without bytes.Buffer
		if err != nil {
			log.Fatal("binary.Write failed")
		}

		//*it's important to find  buf length for readNextBytes and bytes := make([]byte, number)
		//b :=buf.Bytes()
		//l := len(b)
		//fmt.Println(l)
		writeNextBytes(file, buf.Bytes())
	}
	fmt.Println("\n ***** Write File Done ****")
}
func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

// -------------- ReadFile ---------------------------
func readFile() {
	fmt.Println("\nStart to read binary file ------------------------------")
	file, err := os.Open("02.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := payload{}
	for i := 0; i < 10; i++ {
		data := readNextBytes(file, 16)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &m) //binary.LittleEndian
		//err = binary.Read(file, binary.BigEndian, &m) //read file without buffer
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		fmt.Println(m)
	}
	fmt.Println("\n ***** Read File Finished ****")
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// -------------- ReadFile Refactored---------------------------
func readFileRefactored() {
	fmt.Println("\nStart to read binary file Refactored------------------------------")
	file, errOpen := os.Open("02.bin")
	defer file.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	m := payload{}
	for true {
		data, err := readNextBytesRefactored(file, 16)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		buffer := bytes.NewBuffer(*data)
		err = binary.Read(buffer, binary.BigEndian, &m) //binary.LittleEndian
		//err = binary.Read(file, binary.BigEndian, &m) //read file without buffer
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		fmt.Println(m)
	}
	fmt.Println("\n ***** Read File Finished Refactored****")
}

func readNextBytesRefactored(file *os.File, number int) (*[]byte, error) {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
