package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var sums []int64

	for _, line := range fileLines {

		fmt.Println(line)
		if line == "" {
			log.Println("Empty!!!")
			if currentSum > maxSum {
				maxSum = currentSum
				log.Printf("Current sum = %d", currentSum)
			}
			sums = append(sums, currentSum)
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

	fmt.Println(sums)
	// fmt.Println(sort.Sort(sums))
	sort.Slice(sums, func(i, j int) bool { return sums[i] < sums[j] })
	fmt.Println(sums)

	threeLargest := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
	fmt.Println(threeLargest)

	// fmt.Println(fileLines)
	fmt.Printf("The greates sum is %d", maxSum)
}
