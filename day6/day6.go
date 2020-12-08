package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day6\\_in")

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
	set := make(map[string]bool)
	for _, each_ln1 := range text {
		if each_ln1 == "" {
			counter += len(set)
			set = make(map[string]bool)
			continue
		}
		for _, letter := range each_ln1 {
			set[string(letter)] = true
		}
	}
	counter += len(set)

	fmt.Println(counter)
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	numOfPplInGroup := 0
	answers := make(map[string]int)
	for _, each_ln1 := range text {
		if each_ln1 == "" {
			for _, v := range answers {
				if v == numOfPplInGroup {
					counter++
				}
			}
			numOfPplInGroup = 0
			answers = make(map[string]int)
			continue
		}
		for _, letter := range each_ln1 {
			answers[string(letter)]++
		}
		numOfPplInGroup++
	}
	for _, v := range answers {
		if v == numOfPplInGroup {
			counter++
		}
	}

	fmt.Println(counter)
}
