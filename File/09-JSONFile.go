package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	LoggedAt time.Time
}

// https://www.golangprograms.com/golang-writing-struct-to-json-file.html
// https://zetcode.com/golang/json/
func main() {
	EncodeJSONSample()
	DecodeJSONSample()
	JSONPrettyPrint()
	JSONWriteToFile("04.json")
	JSONReadFromFileStructured("04.json")
	JSONReadFromFileUnstructured("04.json")
}
func EncodeJSONSample() {
	fmt.Println("\nEncodeJSONSample ------------------------------")
	//--------------------------------------
	//encoding / Marshalling
	//--------------------------------------
	user0 := User{}
	user0.Id = 1122
	user0.Name = "Peter"
	user0.Password = "Parker"
	user0.LoggedAt = time.Now()

	user0Byte, _ := json.Marshal(user0)
	fmt.Println(string(user0Byte))

	user1 := User{101, "Alice", "wonderland", time.Now()}
	jsonByte, _ := json.Marshal(user1)
	fmt.Println(string(jsonByte))
}

func DecodeJSONSample() {
	fmt.Println("\nDecodeJSONSample ------------------------------")
	//--------------------------------------
	//decoding / Unmarsalling
	//--------------------------------------
	bytes1 := []byte(`{"Id":102,"Name":"Bob","Password":"123xyz","LoggedAt":"2021-10-23T16:08:21.124481-04:00"}`)
	var user2 User
	json.Unmarshal(bytes1, &user2)
	fmt.Println(user2)

	user1 := User{}
	user1.Id = 1122
	user1.Name = "Peter"
	user1.Password = "Parker"
	user1.LoggedAt = time.Now()

	user1Byte, _ := json.Marshal(user1)
	var user4 User
	json.Unmarshal(user1Byte, &user4)
	fmt.Println(user4)
}

func JSONPrettyPrint() {
	fmt.Println("\nJSONPrettyPrint ------------------------------")
	birds := map[string]interface{}{
		"sounds": map[string]string{
			"pigeon":  "coo",
			"eagle":   "squak",
			"owl":     "hoot",
			"duck":    "quack",
			"cuckoo":  "ku-ku",
			"raven":   "cruck-cruck",
			"chicken": "cluck",
			"rooster": "cock-a-doodle-do",
		},
	}

	data, err := json.MarshalIndent(birds, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func JSONWriteToFile(filename string) {
	fmt.Println("\nJSONWriteToFile ------------------------------")
	user := User{}
	user.Id = 1122
	user.Name = "Peter"
	user.Password = "Parker"
	user.LoggedAt = time.Now()
	//...................................
	//Writing struct type to a JSON file
	//...................................
	content, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Reading structured data from JSON Files
func JSONReadFromFileStructured(filename string) {
	fmt.Println("\nJSONReadFromFileStructured ------------------------------")
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	user2 := User{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v\n", user2.Id, user2.Name, user2.Password, user2.LoggedAt)

}

// Reading Unstructured Data from JSON Files
func JSONReadFromFileUnstructured(filename string) {
	fmt.Println("\nJSONReadFromFileUnstructured ------------------------------")
	//...................................
	//Reading into Unstructured type from a JSON file
	//...................................
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var user2 map[string]interface{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}

	// Let's print the unmarshalled data!
	//fmt.Printf("%v\n", user2)
	fmt.Printf("Id: %v\n", user2["Id"])
	fmt.Printf("Name: %s\n", user2["Name"])
	fmt.Printf("Password: %s\n", user2["Password"])
	fmt.Printf("LoggedAt: %v\n", user2["LoggedAt"])
}
