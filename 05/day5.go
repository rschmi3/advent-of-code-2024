package day5

import (
	"bufio"
	"fmt"
	"io/fs"
	"strconv"
	"strings"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

type ruleSet = map[int]Set[int]

// Add inserts an element into the set
func (s Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s Set[T]) Contains(value T) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s.elements)
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	for i := range s2.elements {
		s.Add(i)
	}
	return s
}

func part1(updates [][]int) (total int) {
	for _, v := range updates {
		middle := (len(v) / 2)
		total += v[middle]
	}
	return
}

func part2(fileData []string) (total int) {
	return
}

func generateRules(scanner *bufio.Scanner) (rules ruleSet) {

	rules = make(ruleSet)

	for scanner.Scan() && scanner.Text() != "" {
		line := scanner.Text()
		num1, _ := strconv.Atoi(line[0:strings.Index(line, "|")])
		num2, _ := strconv.Atoi(line[strings.Index(line, "|")+1:])
		// _, num1Found := rules[num1]
		_, num2Found := rules[num2]
		if !num2Found {
			rules[num2] = Set[int]{make(map[int]struct{})}
		}
		rules[num2].Add(num1)
	}
	return
}

func filterUpdates(scanner *bufio.Scanner, rules ruleSet) (updates [][]int) {

outer:
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		update := make([]int, len(line))
		notAllowed := Set[int]{make(map[int]struct{})}

		for i, v := range line {
			num, _ := strconv.Atoi(v)
			if notAllowed.Contains(num) {
				continue outer
			}
			notAllowed = notAllowed.Union(rules[num])
			update[i] = num
		}

		updates = append(updates, update)
	}

	return
}

func Run(file fs.File) {

	scanner := bufio.NewScanner(file)

	rules := generateRules(scanner)
	updates := filterUpdates(scanner, rules)

	part1 := part1(updates)
	part2 := part2(nil)

	fmt.Println(fmt.Sprintf("Day 5 Results: %d, %d", part1, part2))
}