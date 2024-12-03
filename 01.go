package main

import "slices"

func day1Part1(leftList []int, rightList []int) int {

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

func day1Part2(leftList []int, rightList []int) int {
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
