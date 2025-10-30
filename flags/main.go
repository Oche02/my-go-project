package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	inputStr := ""

	// Parse flags and arguments
	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			inputStr = arg
		}
	}

	// Combine input and inserted string
	result := inputStr + insertStr

	// Order if needed
	if orderFlag {
		runes := []rune(result)
		for i := 0; i < len(runes)-1; i++ {
			for j := 0; j < len(runes)-i-1; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}
		result = string(runes)
	}

	// Print final result
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
