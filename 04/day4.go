package day4

import (
	"bufio"
	"fmt"
	"io/fs"
)

func part1(fileData []string) (total int) {

	for i, str := range fileData {
		for j, letter := range str {
			if letter == 'X' {
				leftEnd := j - 3
				rightEnd := j + 4
				upEnd := i - 3
				downEnd := i + 3

				// Left SAMX
				if leftEnd >= 0 {
					candidate := str[leftEnd : j+1]
					if candidate == "SAMX" {
						total++
					}
				}

				// Right XMAS
				if rightEnd <= len(str) {
					candidate := str[j:rightEnd]
					if candidate == "XMAS" {
						total++
					}
				}

				// Up XMAS
				if upEnd >= 0 {
					candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i-1][j], fileData[i-2][j], fileData[i-3][j])
					if candidate == "XMAS" {
						total++
					}

					// Diagonal Left
					if leftEnd >= 0 {
						candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i-1][j-1], fileData[i-2][j-2], fileData[i-3][j-3])
						if candidate == "XMAS" {
							total++
						}
					}

					// Diagonal Right
					if rightEnd <= len(str) {
						candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i-1][j+1], fileData[i-2][j+2], fileData[i-3][j+3])
						if candidate == "XMAS" {
							total++
						}
					}
				}

				// Down XMAS
				if downEnd < len(fileData) {
					candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i+1][j], fileData[i+2][j], fileData[i+3][j])
					if candidate == "XMAS" {
						total++
					}

					// Diagonal Left
					if leftEnd >= 0 {
						candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i+1][j-1], fileData[i+2][j-2], fileData[i+3][j-3])
						if candidate == "XMAS" {
							total++
						}
					}

					// Diagonal Right
					if rightEnd <= len(str) {
						candidate := fmt.Sprintf("%c%c%c%c", fileData[i][j], fileData[i+1][j+1], fileData[i+2][j+2], fileData[i+3][j+3])
						if candidate == "XMAS" {
							total++
						}
					}
				}
			}
		}
	}

	return
}

func part2(fileData []string) (total int) {

	for i, str := range fileData {
		for j, letter := range str {
			if letter == 'A' {
				leftEnd := j - 1
				rightEnd := j + 1
				upEnd := i - 1
				downEnd := i + 1

				if upEnd >= 0 && downEnd < len(fileData) && leftEnd >= 0 && rightEnd < len(str) {
					downRight := fmt.Sprintf("%c%c%c", fileData[upEnd][leftEnd], fileData[i][j], fileData[downEnd][rightEnd])
					downLeft := fmt.Sprintf("%c%c%c", fileData[upEnd][rightEnd], fileData[i][j], fileData[downEnd][leftEnd])
					upRight := fmt.Sprintf("%c%c%c", fileData[downEnd][leftEnd], fileData[i][j], fileData[upEnd][rightEnd])
					upLeft := fmt.Sprintf("%c%c%c", fileData[downEnd][rightEnd], fileData[i][j], fileData[upEnd][leftEnd])

					if (downRight == "MAS" || upLeft == "MAS") && (downLeft == "MAS" || upRight == "MAS") {
						total++
					}
				}

			}
		}
	}
	return
}

func Run(file fs.File) {

	scanner := bufio.NewScanner(file)

	fileData := make([]string, 0)
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}

	part1 := part1(fileData)
	part2 := part2(fileData)

	fmt.Println(fmt.Sprintf("Day 4 Results: %d, %d", part1, part2))
}
