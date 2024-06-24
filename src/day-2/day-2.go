package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// 12 red cubes, 13 green cubes, and 14 blue cubes

func solve(lines []string) {

	const MaxRed = 12
	const MaxGreen = 13
	const MaxBlue = 14

	score := 0 // This is fine, we only ever need one score

	for _, line := range lines {
		splitForID := strings.Split(line, ":")
		// [[Game n], [Other info]

		gameInfo := strings.TrimSpace(splitForID[0])

		gameSplit := strings.Split(gameInfo, " ")
		// [[Game], [n]]
		gameID := gameSplit[1]
		// n

		colors := strings.TrimSpace(splitForID[1])
		colorSegs := strings.Split(colors, ";")
		// [[number of some color, number of some color], [number of some color, number of some color]]

		// We do want to split here because the colors are put back into the bag
		// OK to look at each draw independently

		redCount, greenCount, blueCount := 0, 0, 0
		// This is fine, need to reinit the counts every iteration
		exceeded := false
		// This is also fine, need to reinit exceeded as well

		for _, seg := range colorSegs { // This loop is fine, we need to go through the color segs and break them apart further
			colorsPicked := strings.TrimSpace(seg)
			colorDetails := strings.Split(colorsPicked, ",")
			// [[number of some color], [number of some color]]

			// Logic is fine so far

			for _, detail := range colorDetails { // This is also fine, we can loop over the broken apart strings
				colorNumber := strings.Fields(strings.TrimSpace(detail))
				// fmt.Println(colorNumber) DEBUG
				color := colorNumber[1]
				// fmt.Println(color) DEBUG
				val := colorNumber[0]
				// fmt.Println(val) DEBUG

				value, err := strconv.Atoi(val)

				if err != nil {
					fmt.Println("Issue converting # of color to int")
				}

				// fmt.Println(color)

				switch color {
				case "red":
					redCount += value

					if redCount > MaxRed {
						exceeded = true
						fmt.Println("Invalid game, %d", gameID)
						break // Added break statements, we can get out of here if one of these exceeds
					}

				case "green":
					greenCount += value

					if greenCount > MaxGreen {
						exceeded = true
						fmt.Println("Invalid game, %d", gameID)
						break
					}

				case "blue":
					blueCount += value

					if blueCount > MaxBlue {
						exceeded = true
						fmt.Println("Invalid game, %d", gameID)
						break
					}
				}
			}

			if exceeded {
				break
			}
		}

		if !exceeded {
			game, err := strconv.Atoi(gameID)

			if err != nil {
				fmt.Println("Error converting gameID -> int")
				continue
			}
			score += game
		}

		fmt.Println(score)
	}
}

func main() {

	file, err := os.Open("day-2-data.txt")
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
