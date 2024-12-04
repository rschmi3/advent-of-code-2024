package main

import (
	day1 "advent-of-code/01"
	day2 "advent-of-code/02"
	day3 "advent-of-code/03"
	"embed"
	_ "embed"
	"fmt"
)

//go:embed data/*
var inputs embed.FS

func main() {

	day1Data, err := inputs.Open("data/01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer day1Data.Close()

	day1.Run(day1Data)

	day2Data, err := inputs.Open("data/02.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer day2Data.Close()

	day2.Run(day2Data)

	day3Data, err := inputs.Open("data/03.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer day3Data.Close()

	day3.Run(day3Data)
}
