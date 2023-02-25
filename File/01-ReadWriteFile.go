package main

import (
	"fmt"
	"os"
) 

func main() {
	// simple way to write and read text file
	WriteFile("01.txt")
	ReadFile("01.txt")
}

func WriteFile(filePath string){
    fmt.Println("\nStart to write text file ------------------------------")
    file,err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
		  fmt.Println("Error::Could not open ",filePath)

      return
	  }

	  defer file.Close()
	 
    ln, err2 := file.WriteString("Appending some text to file\n")
	  if err2 != nil {
		  fmt.Println("Error::Could not write text to ",filePath)
	  }else{
      fmt.Println("Success::Operation successful! Text has been appended to ",filePath," ,Length:",ln)
    }	
}

func ReadFile(filePath string){
	fmt.Println("\nStart to read text file ------------------------------")
    file, err := os.ReadFile(filePath)
    if err != nil {
		fmt.Println("Error::Could not open ",filePath)

	return
	}

    fmt.Print(string(file))
}