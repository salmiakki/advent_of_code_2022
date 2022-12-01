package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		text := fileScanner.Text()

		fileLines = append(fileLines, text)
	}

	readFile.Close()

	var currentSum int64 = 0
	var maxSum int64 = 0

	for _, line := range fileLines {

		fmt.Println(line)
		if line == "" {
			log.Println("Empty!!!")
			if currentSum > maxSum {
				maxSum = currentSum
				log.Printf("Current sum = %d", currentSum)
			}
			currentSum = 0
		} else {
			number, _ := strconv.ParseInt(line, 0, 64)
			// fmt.Println(number * 2)
			currentSum += number
		}
	}

	if currentSum > maxSum {
		maxSum = currentSum
	}

	// fmt.Println(fileLines)
	fmt.Printf("The greates sum is %d", maxSum)
}
