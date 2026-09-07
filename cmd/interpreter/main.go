package main

import (
	"fmt"
	"os"
)

func main() {
	// THE RUN CONTRACT: Check if the test harness passed the "--tokenize" flag.
	if len(os.Args) > 1 && os.Args[1] == "--tokenize" {

		// WEEK 1 PROTOTYPE - A simple string to prove single-character logic works.
		source := "(+*-)"
		line := 1

		// SCANNER LOOP - Move character by character using index 'current'
		// Avoid Regular Expressions to build underlying machine from scratch
		for current := 0; current < len(source); current++ {
			char := source[current]

			// CLASSIFICATION - Group raw characters into meaningful token categories
			switch char {
			case '(':
				fmt.Printf("Token(type=LEFT_PAREN, lexeme=(, literal=null, line=%d)\n", line)
			case ')':
				fmt.Printf("Token(type=RIGHT_PAREN, lexeme=), literal=null, line=%d)\n", line)
			case '+':
				fmt.Printf("Token(type=PLUS, lexeme=+, literal=null, line=%d)\n", line)
			case '-':
				fmt.Printf("Token(type=MINUS, lexeme=-, literal=null, line=%d)\n", line)
			case '*':
				fmt.Printf("Token(type=STAR, lexeme=*, literal=null, line=%d)\n", line)
			}
		}

		// EOF TOKEN - Emit an End-Of-File token so the future parser knows when to stop
		fmt.Printf("Token(type=EOF, lexeme=, literal=null, line=%d)\n", line)

		// EXIT CODE: Exit with 0 to indicate a clean scan
		os.Exit(0)
	} else {
		// TIf no flag is passed, default to the interactive loop
		fmt.Println("REPL Mode initialized. Awaiting input...")
	}
}
