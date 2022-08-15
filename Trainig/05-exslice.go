package main

import "fmt"

type part string

func showline(line []part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	assebly := make([]part, 3)
	assebly[0] = "google.com"
	assebly[1] = "you.com"
	assebly[2] = "stackoverflow.com"
	fmt.Println("3 parts:")
	showline(assebly)
	assebly = append(assebly, "w3school", "udemy")
	fmt.Println("added two parts:")
	showline(assebly)
	asseblyslice := assebly[3:]
	fmt.Println("create slice")
	showline(asseblyslice)
}
