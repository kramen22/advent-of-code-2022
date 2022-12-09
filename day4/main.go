package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	start, end int
}

func main() {
	elves := parseElves()

	total := 0
	totalOverlappingAtAll := 0
	for _, pair := range elves {
		if isPairOverlapping(pair[0], pair[1]) {
			total++
		}
		if isPairOverlappingAtAll(pair[0], pair[1]) {
			totalOverlappingAtAll++
		}
	}

	fmt.Printf("got total number of overlapping pairs %d\n", total)
	fmt.Printf("got total number of overlapping at all pairs %d\n", totalOverlappingAtAll)
}

func isPairOverlapping(elf1, elf2 Elf) bool {
	return (elf1.start >= elf2.start && elf1.end <= elf2.end) || (elf2.start >= elf1.start && elf2.end <= elf1.end)
}

func isPairOverlappingAtAll(elf1, elf2 Elf) bool {
	return !(elf1.start > elf2.end || elf2.start > elf1.end)
}

func parseElves() [][]Elf {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := [][]Elf{}

	for fileScanner.Scan() {
		ranges := strings.Split(fileScanner.Text(), ",")

		elfPair := []Elf{}
		for _, zoneRange := range ranges {
			zones := strings.Split(zoneRange, "-")

			zoneStart, _ := strconv.Atoi(zones[0])
			zoneEnd, _ := strconv.Atoi(zones[1])

			elfPair = append(elfPair, Elf{
				start: zoneStart,
				end:   zoneEnd,
			})
		}
		elves = append(elves, elfPair)
	}

	return elves
}
