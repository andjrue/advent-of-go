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

	// const MaxRed = 12 Don't need these for part 2
	// const MaxGreen = 13
	// const MaxBlue = 14

	res := 0 // This is fine, we only ever need one score
	var powerOfMinNeeded = 0

	for _, line := range lines {
		splitForID := strings.Split(line, ":")
		fmt.Println("Split for ID: ", splitForID)
		// [[Game n], [Other info]

		gameInfo := strings.TrimSpace(splitForID[0])
		fmt.Println("gameInfo: ", gameInfo)

		gameSplit := strings.Split(gameInfo, " ")
		fmt.Println("Game Split: ", gameSplit)
		// [[Game], [n]]
		gameID := gameSplit[1]
		fmt.Println("Game ID: ", gameID)
		// n

		colors := strings.TrimSpace(splitForID[1])
		fmt.Println("Colors: ", colors)
		colorSegs := strings.Split(colors, ";")
		fmt.Println("Color Segments: ", colorSegs)
		// [[number of some color, number of some color], [number of some color, number of some color]]

		// We do want to split here because the colors are put back into the bag
		// OK to look at each draw independently

		// exceeded := false
		// This is also fine, need to reinit exceeded as well
		minRed, minGreen, minBlue := 0, 0, 0

		for _, seg := range colorSegs { // This loop is fine, we need to go through the color segs and break them apart further
			colorsPicked := strings.TrimSpace(seg)
			fmt.Println("Colors Picked: ", colorsPicked)
			colorDetails := strings.Split(colorsPicked, ",")
			fmt.Println("Color Details: ", colorDetails)
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

				redCount, greenCount, blueCount := 0, 0, 0
				// This was actually in the wrong place
				// Silly me, we don't want to initialize a count through every big loop. We want it set
				// whenever we start a new iteration through colorDetails.
				// All good, problem solved - solution is working

				// fmt.Println(color)

				switch color {
				case "red":
					redCount += value
					if redCount > minRed {
						minRed = redCount
					}
					fmt.Println("Min Red: ", minRed)

				case "green":
					greenCount += value
					if greenCount > minGreen {
						minGreen = greenCount
					}
					fmt.Println("Min Green: ", minGreen)

				case "blue":
					blueCount += value
					if blueCount > minBlue {
						minBlue = blueCount
					}
					fmt.Println("Min Blue: ", minBlue)
				}
			}
			powerOfMinNeeded = (minBlue * minGreen * minRed)
			fmt.Println(powerOfMinNeeded)
		}
		res += powerOfMinNeeded
		fmt.Println("Result: ", res)
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
