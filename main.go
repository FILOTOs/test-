package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	mp := make(map[string]string)
	mp[".*name.*"] = "My name is Max"
	mp["(hi)|(privet)|(hello)"] = "Hi, nice to meet you"
	mp[".*you.*like.*"] = "I like talking to you :)"

	for {

		fmt.Println("enter your question:")
		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		str = strings.ToLower(strings.Replace(str, "\n", "", -1))

		answerFound := false
		for k, v := range mp {
			if matched, _ := regexp.MatchString(k, str); matched {
				answerFound = true
				fmt.Println(v)

			}

		}
		if !answerFound {
			fmt.Println("I have no idea")
		}
	}
	fmt.Scanln()
}
