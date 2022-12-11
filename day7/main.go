package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sizeMap = make(map[int]*Directory, 0)

type Directory struct {
	Parent   *Directory
	Children map[string]*Directory
	Files    map[string]int
}

func (d *Directory) GetSize() int {
	size := 0

	for key := range d.Children {
		size += d.Children[key].GetSize()
	}

	for key := range d.Files {
		size += d.Files[key]
	}

	sizeMap[size] = d // if theres a collision its whatever i guess :3

	return size
}

func findSusDirs(root *Directory) int {
	total := 0
	for _, dir := range root.Children {
		if size := dir.GetSize(); size < 100000 {
			total += size
		}
		total += findSusDirs(dir)
	}

	return total
}

func main() {
	root := parseDirs()

	// this should populate the size map for part 2
	susSize := findSusDirs(&root)
	fmt.Printf("Got %d total size of dirs < 100000\n", susSize)

	rootSize := root.GetSize()
	sizeLeft := 70000000 - rootSize
	sizeNeeded := 30000000 - sizeLeft
	currentSize := 70000000 // set this high so the first successful one gets set
	fmt.Printf("Root size: %d Size left: %d Size Needed: %d\n", rootSize, sizeLeft, sizeNeeded)
	for size := range sizeMap {
		if size > sizeNeeded && size < currentSize {
			currentSize = size
		}
	}

	fmt.Printf("Got %d size file that is big enough to clear things up :3", currentSize)
}

func parseDirs() Directory {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	root := Directory{
		Children: make(map[string]*Directory, 0),
		Files:    make(map[string]int, 0),
	}
	var current *Directory

	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")

		switch parts[0] {
		case "$": // commands
			switch parts[1] {
			case "ls":
				break // the next parsed line will have the data
			case "cd":
				switch parts[2] {
				case "/":
					current = &root
				case "..":
					current = current.Parent
				default:
					current = current.Children[parts[2]]
				}
			}
		case "dir": // a directory listed by ls
			child := &Directory{
				Parent:   current,
				Children: make(map[string]*Directory, 0),
				Files:    make(map[string]int, 0),
			}

			current.Children[parts[1]] = child
		default: // a file is being listed by ls
			size, _ := strconv.Atoi(parts[0])
			current.Files[parts[1]] = size
		}
	}

	return root
}
