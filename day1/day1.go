package main

import (
	"bufio"
	"fmt" // A package in the Go standard library.
	"regexp"
	"strconv"
	"strings"

	//"io/ioutil" // Implements some I/O utility functions.
	//m "math"    // Math library with local alias m.
	//"net/http"  // Yes, a web server!
	"os" // OS functions like working with the file system
	// String conversions.
)

func main() {
	// Println outputs a line to stdout.
	// It comes from the package fmt.
	//fmt.Println("Hello world!")

	re := regexp.MustCompile(`\d`)

	//////////////////////////////////////////////////////////////

	// Open the file.
	file, err := os.Open("test_input.txt")
	if err != nil {
		// Handle the error and exit.
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after the function ends.

	acc := 0

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Print each line.
		matches := re.FindAllString(line, -1)
		//result := strings.Join(matches, "")
		//fmt.Println(result)

		hit0 := matches[0]
		//hit1 := matches[-1]
		hit1 := matches[len(matches)-1]
		arr := []string{hit0, hit1}
		//inner := strings.Join(arr, "")

		result := strings.Join(arr, "")
		fmt.Println(result)

		//for _, name := range []string{"Bob", "Bill", "Joe"} {
		//fmt.Printf("%s\n", inner)
		//}

		// for _, match := range matches {
		num, _ := strconv.Atoi(result)
		acc += num
		fmt.Printf("%s\n", acc)
		// }

	}

	// Check for errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	////////////////////////////////////////////////////////////////////////

	fmt.Println(acc)
}
