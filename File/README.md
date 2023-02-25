
# File Operations
<br>

## How to append to a file in Golang
Golang allows us to append or add data to an already existing file. We will use its built-in package called os with a method OpenFile() to append text to a file.
```go
os.OpenFile(name/path string, flag int, perm FileMode)
```
### Parameters

**flag**: An int type instruction given to the method to open the file, e.g., read-only, write-only, or read-write. Commonly used flags are as follows:
`O_RDONLY`: It opens the file read-only.
`O_WRONLY`: It opens the file write-only.
`O_RDWR`: It opens the file read-write.
`O_APPEND`: It appends data to the file when writing.
`O_CREATE`: It creates a new file if none exists.
**perm**: A numeric value of the mode that we want os.OpenFile() to execute in, e.g., read-only has a value of 4 and write-only has a value of 2.

### Return values
The os.OpenFile() method returns two values:

**File**: A File on which different operations such as write or append can be performed based on the file mode passed to the function
***PathError**: An error while opening or creating the file.

### Write Text File
```go
file,err := os.OpenFile("sample.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
...
defer file.Close()
...
len, err2 := file.WriteString("Appending some text to file\n")
```
### Read Text File
```go
file, err := os.ReadFile("sample.txt")
...
fmt.Print(string(file))
```