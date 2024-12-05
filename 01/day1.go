package day1

import (
	"bufio"
	"cmp"
	"fmt"
	"io/fs"
	"regexp"
	"slices"
	"strconv"
)

func Insert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func ReadData(scanner bufio.Scanner) ([]int, []int) {

	var leftList []int
	var rightList []int
	for scanner.Scan() {
		line := scanner.Text()

		r, _ := regexp.Compile(`\d+`)
		numbers := r.FindAllString(line, -1)

		leftInt, _ := strconv.Atoi(numbers[0])
		rightInt, _ := strconv.Atoi(numbers[1])
		leftList = Insert(leftList, leftInt)
		rightList = Insert(rightList, rightInt)
	}

	return leftList, rightList
}
func part1(leftList []int, rightList []int) int {

	var totalDifference int

	for i := range len(leftList) {
		difference := leftList[i] - rightList[i]

		if difference < 0 {
			difference = -difference
		}
		totalDifference += (difference)
	}

	return totalDifference
}

func part2(leftList []int, rightList []int) int {
	var totalSimilarity int

	for _, v := range leftList {
		var currentSimilarity int

		idx := slices.Index(rightList, v)
		start := idx + 1

		for idx != -1 {
			currentSimilarity++

			if start >= len(rightList) {
				break
			}

			idx = slices.Index(rightList[start:], v)
			start += idx + 1
		}

		totalSimilarity += currentSimilarity * v
	}
	return totalSimilarity
}

func Run(file fs.File) {

	scanner := bufio.NewScanner(file)

	leftList, rightList := ReadData(*scanner)

	part1 := part1(leftList, rightList)

	part2 := part2(leftList, rightList)
	fmt.Println(fmt.Sprintf("Day 1 Results: %d, %d", part1, part2))
}
