package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getShapeScore(c string) int {
	switch c {
	case "A", "X":
		return ROCK_SCORE
	case "B", "Y":
		return PAPER_SCORE
	case "C", "Z":
		return SCISSORS_SCORE
	default:
		return 0
	}
}

func getOutcomeScore(opponentShape, playerShape string) int {
	opponentShapeScore := getShapeScore(opponentShape)
	playerShapeScore := getShapeScore(playerShape)

	diff := playerShapeScore - opponentShapeScore

	if diff == 1 || diff == -2 {
		return WIN_SCORE
	} else if diff == 0 {
		return DRAW_SCORE
	} else {
		return LOSS_SCORE
	}
}

func getOutcomeScoreFromChar(char string) int {
	switch char {
	case "X":
		return LOSS_SCORE
	case "Y":
		return DRAW_SCORE
	case "Z":
		return WIN_SCORE
	default:
		return 0
	}
}

func getPlayerScore(opponentShape, outcomeChar string) int {
	outcomeScore := getOutcomeScoreFromChar(outcomeChar)
	opponentScore := getShapeScore(opponentShape)
	playerScore := outcomeScore

	if outcomeScore == WIN_SCORE {
		if opponentScore != 3 {
			playerScore += opponentScore + 1
		} else {
			playerScore += 1
		}
	} else if outcomeScore == DRAW_SCORE {
		playerScore += opponentScore
	} else if outcomeScore == LOSS_SCORE {
		if opponentScore == 1 {
			playerScore += 3
		} else {
			playerScore += opponentScore - 1
		}
	}

	return playerScore
}

func partOne(scanner bufio.Scanner) {
	playerScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		shapes := strings.Split(line, " ")
		playerScore += getShapeScore(shapes[1])
		playerScore += getOutcomeScore(shapes[0], shapes[1])
	}

	fmt.Printf("Player score: %d", playerScore)
}

func partTwo(scanner bufio.Scanner) {
	playerScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		shapes := strings.Split(line, " ")
		playerScore += getPlayerScore(shapes[0], shapes[1])
	}

	fmt.Printf("Player score: %d", playerScore)
}

const WIN_SCORE = 6
const DRAW_SCORE = 3
const LOSS_SCORE = 0

const ROCK_SCORE = 1
const PAPER_SCORE = 2
const SCISSORS_SCORE = 3

func main() {
	filePtr := flag.String("file", "demo.txt", "Input file name.")
	partPtr := flag.Int("part", 1, "Problem part. Possible values are [1, 2]")

	flag.Parse()

	fileName := *filePtr
	part := *partPtr

	f, err := os.Open(fileName)
	handleError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if part == 1 {
		partOne(*scanner)
	} else if part == 2 {
		partTwo(*scanner)
	} else {
		log.Fatal("Invalid part specified")
	}
}
