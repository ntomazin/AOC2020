package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day8_in")

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	part1(text)
	part2(text)
}

func part1(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	var counterList []int
	accumulator := 0
	for {
		if Contains(counterList, counter) || counter == len(text) {
			break
		}
		counterList = append(counterList, counter)

		s := strings.Split(text[counter], " ")

		if s[0] == "acc" {
			acc, _ := strconv.Atoi(s[1])
			accumulator += acc
		}

		if s[0] == "jmp" {
			jmp, _ := strconv.Atoi(s[1])
			counter += jmp
			continue
		}
		counter++
	}
	fmt.Println(accumulator)
}

func part2(text []string) {
	//I supposed i needed to change a jmp to nop, not the other way around
	//if needed then lines 85 and 107 should be changed
	var indexesToSkip []int
	for i, line := range text {
		lineParts := strings.Fields(line)
		if lineParts[0] == "jmp" {
			indexesToSkip = append(indexesToSkip, i)
		}
	}

	i := 0
	var accumulator int
	for {
		counter := 0
		var counterList []int
		accumulator = 0
		for {
			if counter >= len(text) {
				break
			}
			if Contains(counterList, counter) {
				i++
				break
			}
			counterList = append(counterList, counter)
			s := strings.Split(text[counter], " ")
			if counter == i {
				s[0] = "nop"
			}
			if s[0] == "acc" {
				acc, _ := strconv.Atoi(s[1])
				accumulator += acc
			}

			if s[0] == "jmp" {
				jmp, _ := strconv.Atoi(s[1])
				counter += jmp
				continue
			}
			counter++

		}
		if counter >= len(text) {
			break
		}
	}
	fmt.Println(accumulator)
}

func Contains(list []int, x int) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}
