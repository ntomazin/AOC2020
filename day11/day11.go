package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day11_in")

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
	var text [][]string

	for scanner.Scan() {
		text = append(text, strings.Split(scanner.Text(), ""))
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	part1(text)
	part2(text)
}

func part1(text [][]string) {
	for {
		changed := false
		seats := make([][]string, len(text))
		for i := range text {
			seats[i] = make([]string, len(text[i]))
			copy(seats[i], text[i])
		}

		for i, _ := range text {
			for j, _ := range text[i] {
				if text[i][j] == "L" && countNeighbors(text, i, j, false) == 0 {
					seats[i][j] = "#"
					changed = true
				} else if text[i][j] == "#" && countNeighbors(text, i, j, false) >= 4 {
					seats[i][j] = "L"
					changed = true
				}
			}
		}
		text = seats

		if !changed {
			break
		}
	}
	output := ""
	for i, _ := range text {
		output = output + strings.Join(text[i], "")
	}
	fmt.Println(strings.Count(output, "#"))
}

func part2(text [][]string) {
	for {
		changed := false
		seats := copyMatrix(text)

		for i, _ := range text {
			for j, _ := range text[i] {
				if text[i][j] == "L" && countNeighbors(text, i, j, true) == 0 {
					seats[i][j] = "#"
					changed = true
				} else if text[i][j] == "#" && countNeighbors(text, i, j, true) >= 5 {
					seats[i][j] = "L"
					changed = true
				}
			}
		}
		text = seats

		if !changed {
			break
		}
	}
	output := ""
	for i, _ := range text {
		output = output + strings.Join(text[i], "")
	}
	fmt.Println(strings.Count(output, "#"))
}

func countNeighbors(seats [][]string, positionX int, positionY int, visible bool) int {
	counterNeighbor := 0
	offset := 0

	if visible {

	}
	positionX = positionX + offset
	positionY = positionY + offset

	if positionX != len(seats)-1 && seats[positionX+1][positionY] == "#" {
		counterNeighbor++
	}
	if positionX != 0 && seats[positionX-1][positionY] == "#" {
		counterNeighbor++
	}
	if positionY != len(seats[0])-1 && seats[positionX][positionY+1] == "#" {
		counterNeighbor++
	}
	if positionY != 0 && seats[positionX][positionY-1] == "#" {
		counterNeighbor++
	}
	if positionX != 0 && positionY != 0 && seats[positionX-1][positionY-1] == "#" {
		counterNeighbor++
	}
	if positionY != 0 && positionX != len(seats)-1 && seats[positionX+1][positionY-1] == "#" {
		counterNeighbor++
	}
	if positionX != 0 && positionY != len(seats[0])-1 && seats[positionX-1][positionY+1] == "#" {
		counterNeighbor++
	}
	if positionX != len(seats)-1 && positionY != len(seats[0])-1 && seats[positionX+1][positionY+1] == "#" {
		counterNeighbor++
	}

	return counterNeighbor
}

func copyMatrix(text [][]string) [][]string {
	seats := make([][]string, len(text))
	for i := range text {
		seats[i] = make([]string, len(text[i]))
		copy(seats[i], text[i])
	}
	return seats
}
