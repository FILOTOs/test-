package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//создаем структуру человека типа "Person"
type Person struct {
	Name string
	Age  int
}

//создаем структуру справочника
type PhoneLookup struct {
	Numbers map[string]Person
}

func NewPhoneLookup() PhoneLookup {
	return PhoneLookup{
		Numbers: map[string]Person{},
	}
}

//метод для добавления номера в справочник
func (p *PhoneLookup) AddNumber(phone, name string, age int) {
	p.Numbers[phone] = Person{
		Name: name,
		Age:  age,
	}
}

//метод где мы получаем номер
func (p *PhoneLookup) GetNumber(phone string) Person {
	if person, ok := p.Numbers[phone]; ok {
		return person
	} else {
		fmt.Println("not found")
		return Person{}
	}
}

func add() error {
	// input phone number
	var phone string
	fmt.Print("input number: ")
	fmt.Scanf("%s\n", &phone)
	// check phone is correct
	if matched, _ := regexp.MatchString("(8|\\+7)[0-9]{10}", phone); !matched {
		return fmt.Errorf("wrong phone")
	}

	fmt.Print("input name: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.ToLower(strings.Replace(name, "\n", "", -1))
	// check phone is correct
	if matched, _ := regexp.MatchString("[a-z]+", name); !matched {
		return fmt.Errorf("wrong name")
	}

	var age int
	fmt.Print("input age: ")
	fmt.Scanf("%d\n", &age)
	if age < 10 {
		return fmt.Errorf("wrong age")
	}

	phoneLookup.AddNumber(phone, name, age)

	return nil
}

func get() {
	var phone string
	fmt.Print("input number: ")
	fmt.Scanf("%s\n", &phone)

	if person := phoneLookup.GetNumber(phone); person.Name != "" {
		fmt.Printf("name: %s, age: %d\n", person.Name, person.Age)
	}
}

var phoneLookup = NewPhoneLookup()

func main() {

	for {
		var cmd string
		fmt.Println("\nInput command: ")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "add":
			if err := add(); err != nil {
				fmt.Println("error:", err.Error())
			}
		case "get":
			get()
		case "exit":
			fmt.Println("bye")
			return
		default:
			fmt.Println("unknown command")
		}
	}

}
