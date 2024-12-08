package main

import (
	day1 "advent-of-code/01"
	day2 "advent-of-code/02"
	day3 "advent-of-code/03"
	day4 "advent-of-code/04"
	day5 "advent-of-code/05"
	day6 "advent-of-code/06"
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
)

//go:embed data/*
var inputs embed.FS

type runFunc func(file fs.File)

func runDay(filename string, fn runFunc) {

	dayData, err := inputs.Open(fmt.Sprintf("data/%s", filename))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer dayData.Close()

	fn(dayData)
}

func main() {

	runDay("01.txt", day1.Run)
	runDay("02.txt", day2.Run)
	runDay("03.txt", day3.Run)
	runDay("04.txt", day4.Run)
	runDay("05.txt", day5.Run)
	runDay("06.txt", day6.Run)
}
