package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	dataByte, err := os.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("\n%s", dataByte)
}

func WriteFile() {
	content := []byte("{\"hello\":\"world\"}")

	err := os.WriteFile("./hello.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteFile() {
	err := os.Remove("./hello.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ReadFile()
	WriteFile()
	// DeleteFile()
}
