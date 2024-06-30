package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(s []string) {
	totalScore := 0
	for _, line := range s {
		getRidOfGame := strings.Split(line, ":")
		// fmt.Println(getRidOfGame)

		removeID := strings.TrimSpace(getRidOfGame[1])
		// fmt.Println(removeID)

		SplitDrawnAndWin := strings.Split(removeID, "|")
		// fmt.Printf("Drawn: %v \n Winners: %v\n", SplitDrawnAndWin[0], SplitDrawnAndWin[1])

		drawnCards := SplitDrawnAndWin[0]
		winningCards := SplitDrawnAndWin[1]

		winners := make(map[int]bool)
		deck := make(map[int]bool)

		winnersList := strings.Fields(winningCards)
		drawnList := strings.Fields(drawnCards)

		for _, drawnStr := range drawnList {
			drawn, err := strconv.Atoi(drawnStr)
			if err != nil {
				fmt.Println("Issue converting drawn # to int: ", err)
				continue
			}

			deck[drawn] = true
		}
		fmt.Println("Drawn Map: ", deck)

		for _, winStr := range winnersList {
			win, err := strconv.Atoi(winStr)
			if err != nil {
				fmt.Println("Error converting winning # to int: ", err)
				continue
			}
			winners[win] = true
		}
		fmt.Println("Winners Map: ", winners)

		score := 0
		for win := range winners {
			if deck[win] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		totalScore += score
	}

	fmt.Println("Final Score: ", totalScore)

}

func main() {

	file, err := os.Open("day-4-data.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	solve(lines)

}
