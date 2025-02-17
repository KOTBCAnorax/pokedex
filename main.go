package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
