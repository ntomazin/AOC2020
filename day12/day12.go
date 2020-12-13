package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day12_in")

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

type Position struct {
	xAxis int
	yAxis int
}

func part1(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	position := 0 + 0i
	direction := 1 + 0i

	for _, each_ln1 := range text {
		command := string(each_ln1[0])
		length, _ := strconv.Atoi(each_ln1[1:len(each_ln1)])
		lengthCmplx := complex(float64(length), 0)

		switch command {
		case "N":
			position += lengthCmplx
		case "E":
			position += lengthCmplx
		case "S":
			position -= lengthCmplx * 1i
		case "W":
			position -= lengthCmplx
		case "L":
			direction *= cmplx.Pow(1i, (lengthCmplx / complex(float64(90), 0)))
		case "R":
			direction /= cmplx.Pow(1i, (lengthCmplx / complex(float64(90), 0)))
		case "F":
			position += lengthCmplx * (direction / complex(cmplx.Abs(direction), 0))
		}

	}
	//fmt.Println(cmplx.Abs(position))
	fmt.Println(math.Abs(real(position)) + math.Abs(imag(position)))
}

func part2(text []string) {
	position := 0 + 0i
	direction := 10 + 1i

	for _, each_ln1 := range text {
		command := string(each_ln1[0])
		length, _ := strconv.Atoi(each_ln1[1:len(each_ln1)])
		lengthCmplx := complex(float64(length), 0)

		switch command {
		case "N":
			position += lengthCmplx
		case "E":
			position += lengthCmplx
		case "S":
			position -= lengthCmplx * 1i
		case "W":
			position -= lengthCmplx
		case "L":
			direction *= cmplx.Pow(1i, (lengthCmplx / complex(float64(90), 0)))
		case "R":
			direction /= cmplx.Pow(1i, (lengthCmplx / complex(float64(90), 0)))
		case "F":
			position += lengthCmplx * direction
		}

	}
	//fmt.Println(cmplx.Abs(position))
	fmt.Println(math.Abs(real(position)) + math.Abs(imag(position)))
}
