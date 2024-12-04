package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sumLine(line string) (sum int) {

	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	mulMatches := mulRegex.FindAllStringSubmatch(line, -1)
	for _, mulMatch := range mulMatches {
		num1, _ := strconv.Atoi(mulMatch[1])
		num2, _ := strconv.Atoi(mulMatch[2])
		sum += num1 * num2
	}
	return sum
}

func calculateSums(scanner bufio.Scanner) (part1Sum int, part2Sum int) {

	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)
	carryDont := false

	for scanner.Scan() {
		line := scanner.Text()

		part1Sum += sumLine(line)

		startValidIdx := 0
		if carryDont {
			startMatch := doRegex.FindStringIndex(line)
			if len(startMatch) > 1 {
				startValidIdx = startMatch[1]
				carryDont = false
			} else {
				startValidIdx = len(line)
			}
		}

		for startValidIdx < len(line) {
			dontMatch := dontRegex.FindStringIndex(line[startValidIdx:])
			endValidIdx := len(line)

			if len(dontMatch) > 0 {
				endValidIdx = startValidIdx + dontMatch[1]
				carryDont = true
			}

			part2Sum += sumLine(line[startValidIdx:endValidIdx])

			doMatch := doRegex.FindStringIndex(line[endValidIdx:])

			if len(doMatch) > 1 {
				startValidIdx = endValidIdx + doMatch[1]
				carryDont = false
			} else {
				break
			}
		}
	}

	return
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

	part1, part2 := calculateSums(*scanner)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
