package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateValidReports(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var validCount int
	for scanner.Scan() {
		line := scanner.Text()

		report := strings.Split(line, " ")

		if isValid(report) {
			validCount++

		} else {
			for i := range report {
				var removed []string
				removed = append(removed, report[0:i]...)
				removed = append(removed, report[i+1:]...)
				if isValid(removed) {
					validCount++
					break
				}
			}
		}

	}

	return validCount
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

	fmt.Println(calculateValidReports(filename))
}
