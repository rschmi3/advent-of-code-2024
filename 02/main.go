package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateValidReports(scanner bufio.Scanner) (part1Count int, part2Count int) {

	for scanner.Scan() {
		line := scanner.Text()

		report := strings.Split(line, " ")

		if isValid(report) {
			part1Count++
			part2Count++

		} else {
			for i := range report {
				var removed []string
				removed = append(removed, report[0:i]...)
				removed = append(removed, report[i+1:]...)
				if isValid(removed) {
					part2Count++
					break
				}
			}
		}
	}

	return
}

func isValid(report []string) bool {
	if len(report) <= 1 {
		return true
	}

	first, _ := strconv.Atoi(report[0])
	second, _ := strconv.Atoi(report[1])
	difference := second - first
	descending := false

	if difference == 0 {
		return false
	} else if difference < 0 {
		if difference < -3 {
			return false
		}
		descending = true
	} else if difference > 3 {
		return false
	}

	first = second
	for _, v := range report[2:] {

		second, _ := strconv.Atoi(v)
		difference := second - first

		if difference == 0 {
			return false
		} else if difference > 0 && (descending || difference > 3) {
			return false
		} else if difference < 0 && (!descending || difference < -3) {
			return false
		}
		first = second
	}
	return true
}

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1, part2 := calculateValidReports(*scanner)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
