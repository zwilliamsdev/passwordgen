package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	length          = 16
	minLength       = 8
	maxLength       = 64
	letters         = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	symbols         = []byte("!@#$&*_+{}[]:?")
	useNumbers bool = true
	useSymbols bool = false
	password   string
)

func printHelp() {
	fmt.Println("Usage: ./passgen <length> <options>")
	fmt.Printf("Default length is %d charaters.\n", length)
	fmt.Println("Options:")
	fmt.Println("  -n: exclude numbers")
	fmt.Println("  -s: include symbols")
	fmt.Println("  -l: specify length (8-64)")
	fmt.Println("Example useage: ./passgen -nsl 32")
	return
}

func processArgs(args []string) {
	// Use default options
	if len(args) == 0 {
		return
	}

	// Check for help flag
	if args[0] == "-h" || args[0] == "--help" {
		printHelp()
		os.Exit(0)
	}

	// Check first argument for flags
	if args[0][0] == '-' {
		for _, flag := range args[0][1:] {
			switch flag {
			case 'n':
				useNumbers = false
			case 's':
				useSymbols = true
			case 'l':
				checkLength(args)
			default:
				fmt.Println("Invalid flag. Use -h or --help for help.")
				os.Exit(1)
			}
		}
	}
}

func checkLength(args []string) {
	if len(args) < 2 {
		fmt.Println("No length specified. Use -h or --help for help.")
		os.Exit(1)
	}

	// Convert string to int
	converted, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid length. Use -h or --help for help.")
		os.Exit(1)
	}

	// Check if length is within range
	if converted < minLength || converted > maxLength {
		fmt.Println("Invalid length. Use -h or --help for help.")
		os.Exit(1)
	}

	length = converted
}

func generatePassword() {
	for i := 0; i < length; i++ {
		// Generate random number between 0 and 2
		// 0 = character, 1 = number, 2 = symbol
		random := rand.Intn(3)

		switch random {
		case 0:
			password += string(generateCharacter())
		case 1:
			if useNumbers {
				password += strconv.Itoa(generateNumber())
			} else {
				i--
			}
		case 2:
			if useSymbols {
				password += string(generateSymbol())
			} else {
				i--
			}
		}
	}
}

func generateCharacter() byte {
	numChars := len(letters)
	char := letters[rand.Intn(numChars)]
	return char
}

func generateNumber() int {
	return rand.Intn(9)
}

func generateSymbol() byte {
	numSymbols := len(symbols)
	symbol := symbols[rand.Intn(numSymbols)]
	return symbol
}

func main() {
	// Arguments from CLI with program stripped
	args := os.Args[1:]
	processArgs(args)
	generatePassword()
	fmt.Printf("Password: %s\nLength: %d\n", password, len(password))
}
