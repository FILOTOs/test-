package main

import "fmt"

func main() {
	for {
		var s string
		fmt.Print("enter your question")
		fmt.Scanf("%s", &s)

		switch s {
		case "name":
			fmt.Println("Max")
		case "age":
			fmt.Printf("I am 24")
		default:
			fmt.Println("What are u talking about?")

		}

		fmt.Scanln()
		fmt.Println()

	}
}
