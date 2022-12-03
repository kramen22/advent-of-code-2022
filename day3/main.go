package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	sacks := parseRucksacks()

	part1(sacks)
	part2(sacks)
}

func part2(sacks []string) {
	errWeight := 0

	for idx := 0; idx+3 <= len(sacks); idx += 3 {
		errWeight += weightMap[findSharedItem(sacks[idx:idx+3]...)]
	}

	fmt.Printf("\nPart 2 final err weight: %d\n", errWeight)
}

func part1(sacks []string) {
	errWeight := 0

	for _, sack := range sacks {
		cmpart1 := sack[0 : len(sack)/2]
		cmpart2 := sack[len(sack)/2:]
		errWeight += weightMap[findSharedItem(cmpart1, cmpart2)]
	}

	fmt.Printf("\nFinal err weight count: %d\n", errWeight)
}

func findSharedItem(ruckSacks ...string) byte {
	sackMaps := []map[byte]interface{}{}
	lastSackIdx := len(ruckSacks) - 1
	for idx := 0; idx < lastSackIdx; idx++ {
		sackMaps = append(sackMaps, mapSack(ruckSacks[idx]))
	}

	lastSack := ruckSacks[lastSackIdx]
	for idx := range lastSack {
		skip := false
		for _, sack := range sackMaps {
			if _, ok := sack[lastSack[idx]]; !ok {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		return lastSack[idx]
	}

	fmt.Printf("Could not find duplicate item!\n")

	return 0
}

func mapSack(sack string) map[byte]interface{} {
	sackMap := map[byte]interface{}{}
	for idx := range sack {
		sackMap[sack[idx]] = struct{}{}
	}

	return sackMap
}

func parseRucksacks() []string {
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

	ruckSacks := []string{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			ruckSacks = append(ruckSacks, record[0])
		}
	}

	return ruckSacks
}

var weightMap = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}
