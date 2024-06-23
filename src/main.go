package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Issue loading env")
	}

	dayOneData := os.Getenv("DAY_1_DATA_PATH")

	var text []string
	file, err := os.Open(dayOneData)

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())

	}

	for i := 0; i < len(text); i++ {
		fmt.Println(text[i])
	}

}
