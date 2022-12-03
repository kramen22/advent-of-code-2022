package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Win      = 6
	Draw     = 3
	Loss     = 0
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

type Move struct {
	you, them int
}

// for part 2
type GameOutcome struct {
	them, outcome int
}

func main() {
	moves, outcomes := parseMoves()

	total := part1(moves)
	totalPart2 := part2(outcomes)

	fmt.Printf("\n\nGot final game value: %d", total)
	fmt.Printf("\n\nGot final game value: %d", totalPart2)
}

func part1(moves []Move) int {
	total := 0
	for _, move := range moves {
		total += gameValue(move)
	}

	return total
}

func part2(outcomes []GameOutcome) int {
	total := 0
	for _, outcome := range outcomes {
		total += gameValueFromOutcome(outcome)
	}

	return total
}

func gameValue(move Move) int {
	switch move.them {
	case move.you:
		return Draw + move.you
	case Rock:
		switch move.you {
		case Paper:
			return Win + move.you
		case Scissors:
			return Loss + move.you
		}
	case Paper:
		switch move.you {
		case Rock:
			return Loss + move.you
		case Scissors:
			return Win + move.you
		}
	case Scissors:
		switch move.you {
		case Rock:
			return Win + move.you
		case Paper:
			return Loss + move.you
		}
	}

	fmt.Printf("Could not play game! Move: %+v\n", move)
	os.Exit(1)

	return -1
}

func gameValueFromOutcome(outcome GameOutcome) int {
	switch outcome.outcome {
	case Draw:
		return Draw + outcome.them
	case Win:
		switch outcome.them {
		case Rock:
			return Win + Paper
		case Paper:
			return Win + Scissors
		case Scissors:
			return Win + Rock
		}
	case Loss:
		switch outcome.them {
		case Rock:
			return Loss + Scissors
		case Paper:
			return Loss + Rock
		case Scissors:
			return Loss + Paper
		}
	}

	fmt.Printf("Could not play game! Move: %+v\n", outcome)
	os.Exit(1)

	return -1
}

func parseMoves() ([]Move, []GameOutcome) {
	filename, ok := os.LookupEnv("CSV_FILE_ENV")
	if !ok {
		filename = "input"
	}

	// Open the file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not read file! err: %s", err.Error())

		os.Exit(1)
	}

	defer csvFile.Close()

	// Parse the file
	r := csv.NewReader(csvFile)

	moves := []Move{}
	outcomes := []GameOutcome{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// Hack but whatever
		parts := strings.Split(record[0], " ")

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			moves = append(moves, Move{
				them: mapMove(parts[0]),
				you:  mapMove(parts[1]),
			})
			outcomes = append(outcomes, GameOutcome{
				them:    mapMove(parts[0]),
				outcome: mapOutcome(parts[1]),
			})
		}
	}

	return moves, outcomes
}

func mapMove(encryptedMove string) int {
	switch encryptedMove {
	case "X", "A":
		return Rock
	case "Y", "B":
		return Paper
	case "Z", "C":
		return Scissors
	}

	fmt.Printf("Unexpected move: %s\n", encryptedMove)
	os.Exit(1)

	return -1
}

func mapOutcome(encryptedOutcome string) int {
	switch encryptedOutcome {
	case "X":
		return Loss
	case "Y":
		return Draw
	case "Z":
		return Win
	}

	fmt.Printf("Unexpected move: %s\n", encryptedOutcome)
	os.Exit(1)

	return -1
}
