package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func splitStrings(l []string) int {
	total := 0
	re := regexp.MustCompile("[0-9]") //

	for _, line := range l {
		digits := re.FindAllString(line, -1)

		if len(digits) == 1 { // Need to check if only one digit found in string, answer was wrong before adding this
			firstDig := digits[0]
			secDig := digits[0]
			num := firstDig + secDig

			i, err := strconv.Atoi(num) // https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go

			if err != nil {
				fmt.Print("Issue converting to int")
			}

			total += i
		}

		if len(digits) >= 2 {
			firstDig := digits[0]
			secDig := digits[len(digits)-1]
			num := firstDig + secDig

			i, err := strconv.Atoi(num)

			if err != nil {
				fmt.Print("Issue converting to int")
			}

			total += i
		}
	}

	// I feel like there is a better way than repeating myself, but this works

	return total
}

func main() {

	file, err := os.Open("day-1-input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(splitStrings(lines))

}

// Returning the correct answer
// This has been a lot of fun to do in Go
// Googling is required
