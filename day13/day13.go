package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day13_in")

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
	minutes, _ := strconv.Atoi(text[0])
	timestamps := strings.Split(text[1], ",")
	minWait := 9999999
	bestId := -1

	for _, timestamp := range timestamps {
		if timestamp == "x" {
			continue
		}
		id, _ := strconv.Atoi(timestamp)

		waitTime := int(math.Floor(float64(minutes/id)) + 1)
		if minWait > ((waitTime * id) % minutes) {
			bestId = id
			minWait = ((waitTime * id) % minutes)
		}
	}
	fmt.Println(bestId * minWait)

}

func part2(text []string) {
	timestamps := strings.Split(text[1], ",")
	start := 0
	step := 1

	for i, s := range timestamps {
		bus, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		for (start+i)%bus != 0 {
			start += step
		}
		step *= bus
	}
	fmt.Println(start)
}
