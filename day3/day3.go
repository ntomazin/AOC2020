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
	file, err := os.Open("day3_in")

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
	counter := 0
	startingIndex := 0
	for _, each_ln1 := range text {
		if string(each_ln1[startingIndex]) == "#" {
			counter++
		}
		startingIndex += 3
		startingIndex = startingIndex % len(each_ln1)
	}
	fmt.Println(counter)
}

func part2(text []string) {
	counter_all := 1
	for _, i := range []int{1, 3, 5, 7} {
		counter := 0
		startingIndex := 0

		for _, each_ln1 := range text {
			if string(each_ln1[startingIndex]) == "#" {
				counter++
			}
			startingIndex += i
			startingIndex = startingIndex % len(each_ln1)
		}
		counter_all *= counter
	}
	counter := 0
	startingIndex := 0

	for i, each_ln1 := range text {
		if i%2 == 1 {
			continue
		}
		if string(each_ln1[startingIndex]) == "#" {
			counter++
		}
		startingIndex += 1
		startingIndex = startingIndex % len(each_ln1)
	}
	counter_all *= counter

	fmt.Println(counter_all)
}
