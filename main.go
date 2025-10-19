package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0;; i++ {
		fmt.Print("Pokedex > ")
		var input string 

		if scanner.Scan() {
			input = scanner.Text()
		}

		firstWord := cleanInput(input)[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

