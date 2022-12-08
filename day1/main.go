package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	fruit         []int
	totalCalories int
}

type ByTotalCalories []Elf

func (s ByTotalCalories) Len() int {
	return len(s)
}
func (s ByTotalCalories) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByTotalCalories) Less(i, j int) bool {
	return s[i].totalCalories < s[j].totalCalories
}

func main() {
	elves := parseElves()

	sort.Sort(ByTotalCalories(elves))

	fmt.Printf("Got most calories for elf: %+v\n", elves[len(elves)-1])

	// for part 2.
	topThree := 0
	for i := 0; i < 3; i++ {
		topThree += elves[len(elves)-1-i].totalCalories
	}

	fmt.Printf("Got top three total calories of: %d", topThree)
}

func (e *Elf) addFruit(fruit int) {
	e.fruit = append(e.fruit, fruit)
	e.totalCalories += fruit
}

func parseElves() []Elf {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elves := []Elf{}
	toAdd := Elf{}

	for fileScanner.Scan() {
		switch len(fileScanner.Text()) {
		case 0:
			elves = append(elves, toAdd)
			toAdd = Elf{}
		default:
			fruitAmt, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Printf("error parsing: %s\n", err.Error())

				os.Exit(1)
			}
			toAdd.addFruit(fruitAmt)
		}
	}

	return elves
}
