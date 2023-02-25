package main

import (
	"fmt"
	"io"
	"os"
)

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := os.Open("01.txt")
    checkErr(err)

    w1 := make([]byte, 9)
    ln1, err := f.Read(w1)
    checkErr(err)
    fmt.Printf("Read bytes:%d  , word1:%s\n", ln1, string(w1[:ln1]))

	/*
    Seek sets the offset for the next Read or Write on file to offset, 
	interpreted according to whence: 
	0 means relative to the origin of the file, 
	1 means relative to the current offset, 
    2 means relative to the end. 
	It returns the new offset and an error, if any. 
	The behavior of Seek on a file opened with O_APPEND is not specified. 
	If f is a directory, the behavior of Seek varies by operating system; 
	you can seek to the beginning of the directory on Unix-like operating systems, 
	but not on Windows.
	io.SeekStart = 0: the file's starting point is used.
    io.SeekCurrent = 1: the seek is made in relation to the file's most recent offset.
	io.SeekEnd = 2: it searches in relation to the file's end.
	*/
	start, err := f.Seek(10, 0) //io.SeekStart
    checkErr(err)
    w2 := make([]byte, 4)
    ln2, err := f.Read(w2)
    checkErr(err)
    fmt.Printf("Read bytes:%d ,  From:%d , word2:%s\n", ln2, start, string(w2[:ln2]))

	end, err := f.Seek(-5, io.SeekEnd) //io.SeekEnd
    checkErr(err)
    w3 := make([]byte, 4)
    ln3, err := f.Read(w3)
    checkErr(err)
    fmt.Printf("Read bytes:%d ,  From:%d , word3:%s\n", ln3, end, string(w3[:ln3]))

    f.Close()
}