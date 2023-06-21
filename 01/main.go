package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sum(arr []int) int {
	sum := 0
	for idx, _ := range arr {
		sum += arr[idx]
	}
	return sum
}

func partOne(scanner bufio.Scanner) {
	maxCalories := 0
	calories := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
			continue
		}
		calorie, err := strconv.Atoi(scanner.Text())
		handleError(err)
		calories += calorie

	}

	fmt.Printf("Highest calories: %d", maxCalories)
}

func partTwo(scanner bufio.Scanner) {
	LIMIT := 3

	maxCalories := make([]int, LIMIT)
	calories := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			for i, _ := range maxCalories {
				if calories > maxCalories[i] {
					// swap
					t := maxCalories[i]
					maxCalories[i] = calories
					calories = t
				}
			}

			calories = 0
			continue
		}
		calorie, err := strconv.Atoi(scanner.Text())
		handleError(err)
		calories += calorie
	}

	// for last iteration
	for i, _ := range maxCalories {
		if calories > maxCalories[i] {
			// swap
			t := maxCalories[i]
			maxCalories[i] = calories
			calories = t
		}
	}

	fmt.Printf("Sum of %d highest calories: %d", LIMIT, sum(maxCalories))

}

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
