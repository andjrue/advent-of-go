package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s []string) {
	numberStart := -1
	numberEnd := -1

	for row, line := range s {

		for column, r := range line {

			if r >= 0 && r <= 9 {
				if numberStart == -1 {
					numberStart = column
					continue
				}

				numberEnd = column
			}
		}
	}
}

func main() {

	file, err := os.Open("test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//for i := 0; i < len(lines); i++ {
	//	fmt.Println(lines[i])
	//}

	solve(lines)

}
