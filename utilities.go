package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func Insert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func ReadFile(filename string) ([]int, []int, error) {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return make([]int, 0), make([]int, 0), err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	return leftList, rightList, nil
}
