package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day15_in")

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

	startTime := time.Now()
	part1(text)
	fmt.Println("Part 1 took:", time.Since(startTime))

	startTime = time.Now()
	part2(text)
	fmt.Println("Part 2 took:", time.Since(startTime))
}

func part1(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	spoken, last := map[int]int{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(text[0])), ",") {
		last, _ = strconv.Atoi(s)
		spoken[last] = i + 1
	}

	for i := len(spoken); i < 2020; i++ {
		if v, ok := spoken[last]; ok {
			spoken[last], last = i, i-v
		} else {
			spoken[last], last = i, 0
		}
	}
	fmt.Println(last)
}

func part2(text []string) {
	spoken, last := map[int]int{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(text[0])), ",") {
		last, _ = strconv.Atoi(s)
		spoken[last] = i + 1
	}

	for i := len(spoken); i < 30000000; i++ {
		if v, ok := spoken[last]; ok {
			spoken[last], last = i, i-v
		} else {
			spoken[last], last = i, 0
		}
	}
	fmt.Println(last)
}
