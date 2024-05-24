package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Persons []Person

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) getGender() string {
	return p.Gender
}

func (p *Person) canVoteOrNot() bool {
	return p.Age >= 18
}

func (persons *Persons) addPerson(name string, age int, gender string) {
	*persons = append(*persons, Person{
		Name:   name,
		Age:    age,
		Gender: gender,
	})
}

func (persons *Persons) listAllPerson() {
	for _, person := range *persons {
		fmt.Printf("%+v\n", person)
	}
}

func main() {
	person := Person{
		Name:   "Alisa",
		Age:    20,
		Gender: "female",
	}

	var persons Persons
	persons.addPerson("Sahas", 24, "male")
	persons.listAllPerson()

	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// input = strings.TrimSpace(input)

	fmt.Println("name", person.getName())
	fmt.Println("gender", person.getGender())
	fmt.Println("can vote or not", person.canVoteOrNot())
}
