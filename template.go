package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func partOne(scanner bufio.Scanner) {
}

func partTwo(scanner bufio.Scanner) {
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
