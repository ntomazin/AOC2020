package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day10_in")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []int
	for scanner.Scan() {
		v := scanner.Text()
		in, _ := strconv.Atoi(v)
		input = append(input, in)
	}

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []int) {
	sort.Ints(input)
	counter1 := 1
	counter3 := 1

	for i, _ := range input {
		if i == 0 {
			continue
		}
		if input[i]-input[i-1] == 1 {
			counter1++
		} else if input[i]-input[i-1] == 3 {
			counter3++
		}
	}
	fmt.Println(counter1 * counter3)
}

func part2(input []int) {
	sort.Ints(input)
	permutation := map[int]int{0: 1}

	for _, i := range input {
		permutation[i] = permutation[i-1] + permutation[i-2] + permutation[i-3]
	}

	fmt.Println(permutation[input[len(input)-1]])

}
