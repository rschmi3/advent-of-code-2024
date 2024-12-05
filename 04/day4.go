package day4

import (
	"bufio"
	"fmt"
	"io/fs"
)

func part1(fileData []string) (total int) {

	for i, v := range fileData {
		for j := range len(v) {
			if v[j] == 'X' {
				leftEnd := j - 3
				rightEnd := j + 4
				upEnd := i - 3
				downEnd := i + 3

				// Left SAMX
				if leftEnd >= 0 {
					candidate := v[leftEnd : j+1]
					if candidate == "SAMX" {
						total++
					}
				}

				// Right XMAS
				if rightEnd <= len(v) {
					candidate := v[j:rightEnd]
					if candidate == "XMAS" {
						total++
					}
				}

				// Up XMAS
				if upEnd >= 0 {
					candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i-1][j]), string(fileData[i-2][j]), string(fileData[i-3][j]))
					if candidate == "XMAS" {
						total++
					}

					// Diagonal Left
					if leftEnd >= 0 {
						candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i-1][j-1]), string(fileData[i-2][j-2]), string(fileData[i-3][j-3]))
						if candidate == "XMAS" {
							total++
						}
					}

					// Diagonal Right
					if rightEnd <= len(v) {
						candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i-1][j+1]), string(fileData[i-2][j+2]), string(fileData[i-3][j+3]))
						if candidate == "XMAS" {
							total++
						}
					}
				}

				// Down XMAS
				if downEnd < len(fileData) {
					candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i+1][j]), string(fileData[i+2][j]), string(fileData[i+3][j]))
					if candidate == "XMAS" {
						total++
					}

					// Diagonal Left
					if leftEnd >= 0 {
						candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i+1][j-1]), string(fileData[i+2][j-2]), string(fileData[i+3][j-3]))
						if candidate == "XMAS" {
							total++
						}
					}

					// Diagonal Right
					if rightEnd <= len(v) {
						candidate := fmt.Sprintf("%s%s%s%s", string(fileData[i][j]), string(fileData[i+1][j+1]), string(fileData[i+2][j+2]), string(fileData[i+3][j+3]))
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

func Run(file fs.File) {

	scanner := bufio.NewScanner(file)

	fileData := make([]string, 0)
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}

	part1 := part1(fileData)

	fmt.Println(fmt.Sprintf("Day 2 Results: %d", part1))
}
