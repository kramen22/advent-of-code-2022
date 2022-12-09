//                [M]     [V]     [L]
//[G]             [V] [C] [G]     [D]
//[J]             [Q] [W] [Z] [C] [J]
//[W]         [W] [G] [V] [D] [G] [C]
//[R]     [G] [N] [B] [D] [C] [M] [W]
//[F] [M] [H] [C] [S] [T] [N] [N] [N]
//[T] [W] [N] [R] [F] [R] [B] [J] [P]
//[Z] [G] [J] [J] [W] [S] [H] [S] [G]
// 1   2   3   4   5   6   7   8   9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// im lazy.
var crates = []ElfStack{
	{
		stack: []string{"Z", "T", "F", "R", "W", "J", "G"},
	},
	{
		stack: []string{"G", "W", "M"},
	},
	{
		stack: []string{"J", "N", "H", "G"},
	},
	{
		stack: []string{"J", "R", "C", "N", "W"},
	},
	{
		stack: []string{"W", "F", "S", "B", "G", "Q", "V", "M"},
	},
	{
		stack: []string{"S", "R", "T", "D", "V", "W", "C"},
	},
	{
		stack: []string{"H", "B", "N", "C", "D", "Z", "G", "V"},
	},
	{
		stack: []string{"S", "J", "N", "M", "G", "C"},
	},
	{
		stack: []string{"G", "P", "N", "W", "C", "J", "D", "L"},
	},
}

type Move struct {
	amount, from, to int
}

type ElfStack struct {
	stack []string
}

func (es *ElfStack) Pop() string {
	tmp := es.stack[len(es.stack)-1]
	es.stack = es.stack[0 : len(es.stack)-1]

	return tmp
}

func (es *ElfStack) PopV2(amount int) []string {
	tmp := es.stack[len(es.stack)-amount : len(es.stack)]
	es.stack = es.stack[0 : len(es.stack)-amount]

	return tmp
}

func (es *ElfStack) Push(crate ...string) {
	es.stack = append(es.stack, crate...)
}

func (stack *ElfStack) Move(amount int, to *ElfStack) {
	for idx := 0; idx < amount; idx++ {
		to.Push(stack.Pop())
	}
}

func (stack *ElfStack) MoveV2(amount int, to *ElfStack) {
	to.Push(stack.PopV2(amount)...)
}

func main() {
	moves := parseMoves()
	cratesPart2 := []ElfStack{}
	for _, crate := range crates {
		cratesPart2 = append(cratesPart2, ElfStack{stack: crate.stack})
	}

	for _, move := range moves {
		crates[move.from-1].Move(move.amount, &crates[move.to-1])
		cratesPart2[move.from-1].MoveV2(move.amount, &cratesPart2[move.to-1])
	}

	for _, crate := range crates {
		fmt.Printf("%s", crate.Pop())
	}

	fmt.Println()

	for _, crate := range cratesPart2 {
		fmt.Printf("%s", crate.Pop())
	}
}

func parseMoves() []Move {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	moves := []Move{}

	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")

		amount, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		moves = append(moves, Move{
			amount: amount,
			to:     to,
			from:   from,
		})
	}

	return moves
}
