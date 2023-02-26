package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ioutil.WriteFile and ioutil.ReadFile is deprecated for go > 1.16
func main() {

	//! Go 1.16 or later :: Replace ioutil with os
	mydata := []byte("Appending some text to file\n")

	//------------- Write to file
	// the WriteFile method returns an error if unsuccessful
	//* Delete 03.txt file if exist and re-create it
	fmt.Println("\nStart to write text file ------------------------------")
	err := ioutil.WriteFile("03.txt", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	//------------- Read from file
	fmt.Println("\nStart to read text file ------------------------------")
	data, err := ioutil.ReadFile("03.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

	//------------- Write to file
	// the WriteFile method returns an error if unsuccessful
	fmt.Println("\nStart to write text file ------------------------------")
	f, err := os.OpenFile("03.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new Appending some text to file\n"); err != nil {
		panic(err)
	}

	//------------- Read from file
	fmt.Println("\nStart to read text file ------------------------------")
	data, err = ioutil.ReadFile("03.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}
