package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		inputs := cleanInput(text)
		command := inputs[0]

		fmt.Printf("Your command was: %s\n", command)

	}
}
