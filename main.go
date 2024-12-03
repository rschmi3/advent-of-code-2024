package main

import (
	"fmt"
	"os"
)

func main() {

	filename := os.Args[1]

	leftList, rightList, err := ReadFile(filename)

	if err != nil {
		return
	}

	day1Part1Result := day1Part1(leftList, rightList)

	day1Part2Result := day1Part2(leftList, rightList)
	fmt.Println(fmt.Sprintf("Day 1 Results: %d, %d", day1Part1Result, day1Part2Result))
}
