package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var movePointsMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var translationMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var drawPoints = 3
var winPoints = 6
var losePoints = 0

func calculatePoints(villain string, me string) int {
	movePoints := movePointsMap[me]
	log.Printf("movePoints = %d", movePoints)
	// log.Printf("movePoints + drawPoints = %d", movePoints+drawPoints)
	if translationMap[villain] == me {
		return (movePoints + drawPoints)
	}
	if (villain == "A" && me == "Z") || (villain == "B" && me == "X") || (villain == "C" && me == "Y") {
		return movePoints
	}
	return movePoints + winPoints
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	points := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		s := strings.Fields(line)
		roundPoints := calculatePoints(s[0], s[1])
		log.Println(roundPoints)
		points += roundPoints
	}
	fmt.Println(points)

	readFile.Close()
}
