package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		// Perform reverse operation if there are no additional arguments
		reverse(args)
	} else {
		fmt.Println("Too many arguments")
	}
}

// Define a flag to specify the input file for the reverse operation
var readFlags = flag.String("reverse", "example.txt", "read file from flag")

// Reverse function handles the input checks and performs the reverse operation.
func reverse(args []string) {
	// Check for potential usage errors
	checkForAudit()

	// Define the path to the ASCII fonts file
	fonts := "standard.txt"

	// Define usage instructions
	const usage = "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>"

	// Check if the --reverse flag is provided with a single argument
	if !strings.Contains(*readFlags, "--reverse=") && len(args) == 1 {
		fmt.Println(args[0])
		os.Exit(0)
	}

	// Check for additional arguments
	if len(args) > 0 {
		fmt.Println(usage)
		return
	}

	// Read the content of the input file
	input, err := os.ReadFile("examples/" + *readFlags)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
	}

	// Split the input into lines
	matrix := strings.Split(string(input), "\n")

	// Remove trailing dollar signs from each line
	matrix2 := delDollarSigns(matrix)

	// Find empty columns (spaces) in the user input
	spaces := findSpace(matrix2)

	// Split the user input based on empty columns
	userInput := splitUserInput(matrix2, spaces)

	// Map the split user input
	userInputMap := userInputMapping(userInput)

	// Get the ASCII graphic fonts
	asciiGraphic := getASCIIgraphicFont(fonts)

	// Match user input with ASCII graphic fonts and generate the output
	output := mapUserInputWithASCIIgraphicFont(userInputMap, asciiGraphic)

	// Print the generated output
	fmt.Println(output)
}

// delDollarSigns removes trailing dollar signs from each line in the matrix.
func delDollarSigns(matrix []string) []string {
	var matrix2 []string
	for _, v := range matrix {
		lenv := len(v)
		if lenv <= 1 {
			matrix2 = append(matrix2, "")
		} else {
			matrix2 = append(matrix2, v[:lenv-1])
		}
	}
	return matrix2
}

// findSpace identifies empty columns in the matrix.
func findSpace(matrix []string) []int {
	var emptyColumns []int
	count := 0

	for column := 0; column < len(matrix[0]); column++ {
		for row := 0; row < len(matrix)-1; row++ {
			if matrix[row][column] == ' ' {
				count++
			} else {
				count = 0
				break
			}

			if count == len(matrix)-1 {
				emptyColumns = append(emptyColumns, column)
				count = 0
			}
		}
	}

	// Check for extra spaces and convert them accordingly
	count = 5
	var indexToRem []int
	for i := range emptyColumns {
		if count == 0 {
			count = 5
			continue
		}
		if i > 0 {
			if emptyColumns[i] == (emptyColumns[i-1]) + 1 {
				indexToRem = append(indexToRem, i)
				count -= 1
			}
		}
	}

	// Remove extra spaces
	for i := len(indexToRem) - 1; i >= 0; i-- {
		emptyColumns = removeIndex(emptyColumns, indexToRem[i])
	}

	return emptyColumns
}

// removeIndex removes an element at a given index from a slice.
func removeIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// checkForAudit checks for potential usage errors.
func checkForAudit() {
	if strings.Contains(os.Args[1], "--") && !strings.Contains(os.Args[1], "=") {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		os.Exit(0)
	}
}

// splitUserInput splits the user input based on empty columns.
func splitUserInput(matrix []string, emptyColumns []int) string {
	var result string
	result = "\n"
	start := 0
	end := 0

	for _, column := range emptyColumns {
		if end < len(matrix[0]) {
			end = column

			for _, characters := range matrix {
				if len(characters) > 0 {
					columns := characters[start:end]
					result = result + columns + " "
				}
				result = result + "\n"
			}

			start = end + 1
		}
	}

	return result
}

// userInputMapping maps the user input for search.
func userInputMapping(result string) map[int][]string {
	strSlice := strings.Split(result, "\n")
	graphicInput := make(map[int][]string)
	j := 0

	for _, ch := range strSlice {
		if ch == "" {
			j++
		} else {
			graphicInput[j] = append(graphicInput[j], ch)
		}
	}

	return graphicInput
}

// getASCIIgraphicFont reads ASCII graphic fonts from a file.
func getASCIIgraphicFont(fonts string) map[int][]string {
	readFile, err := os.ReadFile(fonts)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
	}

	slice := strings.Split(string(readFile), "\n")
	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}

	return ascii
}

// mapUserInputWithASCIIgraphicFont matches user input with ASCII graphic fonts and returns the output.
func mapUserInputWithASCIIgraphicFont(graphicInput, ascii map[int][]string) string {
	var keys []int
	for k := range graphicInput {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	var output string
	var sliceOfBytes []byte

	for _, value := range keys {
		graphicValue := graphicInput[value]
		for asciiKey, asciiValue := range ascii {
			if reflect.DeepEqual(asciiValue, graphicValue) {
				sliceOfBytes = append(sliceOfBytes, byte(asciiKey))
			}
		}
		output = string(sliceOfBytes)
	}

	return output
}
