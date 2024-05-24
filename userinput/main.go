package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) (string, error) {
	prompt = strings.TrimSpace(prompt)
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	fmt.Println("Welcome to user input")

	name, err := getInput("What is your name?")
	printError(err)
	fmt.Println("My Name is", name)

	age, err := getInput("What is your Age?")
	printError(err)
	ageInNumber, err := strconv.ParseInt(age, 10, 16)
	fmt.Printf("My age is %d", ageInNumber)

}
