package day6

import (
	"bufio"
	"fmt"
	"io/fs"
)

func part1(lab [][]rune, startIndex [2]int) (total int) {

	currentIndex := startIndex
	currentDirection := lab[startIndex[0]][startIndex[1]]
	lab[currentIndex[0]][currentIndex[1]] = 'X'
	total = 1

	for {
		if lab[currentIndex[0]][currentIndex[1]] == '.' {
			total++
			lab[currentIndex[0]][currentIndex[1]] = 'X'
		}

		var nextIndex [2]int
		if currentDirection == '^' {
			nextIndex[0] = currentIndex[0] - 1
			nextIndex[1] = currentIndex[1]
			if nextIndex[0] < 0 {
				break
			}

			if lab[nextIndex[0]][nextIndex[1]] == '#' {
				currentDirection = '>'
				nextIndex = currentIndex
			}
		} else if currentDirection == '>' {
			nextIndex[0] = currentIndex[0]
			nextIndex[1] = currentIndex[1] + 1
			if nextIndex[1] >= len(lab[currentIndex[0]]) {
				break
			}

			if lab[nextIndex[0]][nextIndex[1]] == '#' {
				currentDirection = 'v'
				nextIndex = currentIndex
			}
		} else if currentDirection == 'v' {
			nextIndex[0] = currentIndex[0] + 1
			nextIndex[1] = currentIndex[1]
			if nextIndex[0] >= len(lab[currentIndex[1]]) {
				break
			}

			if lab[nextIndex[0]][nextIndex[1]] == '#' {
				currentDirection = '<'
				nextIndex = currentIndex
			}
		} else if currentDirection == '<' {
			nextIndex[0] = currentIndex[0]
			nextIndex[1] = currentIndex[1] - 1
			if nextIndex[1] < 0 {
				break
			}

			if lab[nextIndex[0]][nextIndex[1]] == '#' {
				currentDirection = '^'
				nextIndex = currentIndex
			}
		}
		currentIndex = nextIndex
	}

	return
}
func part2() (total int) {
	return
}

func readFile(file fs.File) (lab [][]rune, startIndex [2]int) {
	scanner := bufio.NewScanner(file)

	lab = make([][]rune, 0)
	for i := 0; scanner.Scan(); i++ {
		lab = append(lab, make([]rune, len(scanner.Text())))
		for j, letter := range scanner.Text() {
			if letter == '^' || letter == '>' || letter == '<' || letter == 'v' {
				startIndex[0] = i
				startIndex[1] = j
			}
			lab[i][j] = letter
		}
	}
	return
}

func Run(file fs.File) {

	lab, startIndex := readFile(file)

	part1 := part1(lab, startIndex)
	part2 := part2()

	fmt.Println(fmt.Sprintf("Day 5 Results: %d, %d", part1, part2))
}
