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

// Create new set of type T
func NewSet[T comparable]() Set[T] {
	return Set[T]{elements: make(map[T]struct{})}
}

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

func part2(updates [][]int, rules ruleSet) (total int) {
	for _, update := range updates {

		for i := range update {
			for j := len(update) - 1; j > i; j-- {
				rule, ok := rules[update[i]]
				if ok && rule.Contains(update[j]) {
					temp := update[i]
					update[i] = update[j]
					update[j] = temp
				}
			}

		}

		middle := (len(update) / 2)
		total += update[middle]
	}
	return
}

func generateRules(scanner *bufio.Scanner) (rules ruleSet) {

	rules = make(ruleSet)

	for scanner.Scan() && scanner.Text() != "" {
		line := scanner.Text()
		num1, _ := strconv.Atoi(line[0:strings.Index(line, "|")])
		num2, _ := strconv.Atoi(line[strings.Index(line, "|")+1:])
		_, num2Found := rules[num2]
		if !num2Found {
			rules[num2] = NewSet[int]()
		}
		rules[num2].Add(num1)
	}
	return
}

func generateUpdates(scanner *bufio.Scanner) (updates [][]int) {

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		update := make([]int, len(line))

		for i, v := range line {
			num, _ := strconv.Atoi(v)
			update[i] = num
		}

		updates = append(updates, update)
	}
	return
}

func filterUpdates(updates [][]int, rules ruleSet) (validUpdates [][]int, invalidUpdates [][]int) {

outer:
	for _, update := range updates {

		notAllowed := NewSet[int]()
		for _, num := range update {

			if notAllowed.Contains(num) {
				invalidUpdates = append(invalidUpdates, update)
				continue outer
			}
			notAllowed = notAllowed.Union(rules[num])
		}
		validUpdates = append(validUpdates, update)
	}

	return
}

func Run(file fs.File) {

	scanner := bufio.NewScanner(file)

	rules := generateRules(scanner)
	updates := generateUpdates(scanner)
	validUpdates, invalidUpdates := filterUpdates(updates, rules)

	part1 := part1(validUpdates)
	part2 := part2(invalidUpdates, rules)

	fmt.Println(fmt.Sprintf("Day 5 Results: %d, %d", part1, part2))
}
