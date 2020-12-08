package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day1_in")

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
	for _, each_ln1 := range text {
		for _, each_ln2 := range text {
			var value1, _ = strconv.Atoi(each_ln1)
			var value2, _ = strconv.Atoi(each_ln2)

			if value1+value2 == 2020 {
				fmt.Println(value1 * value2)
			}

		}
	}
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	for _, each_ln1 := range text {
		for _, each_ln2 := range text {
			for _, each_ln3 := range text {
				var value1, _ = strconv.Atoi(each_ln1)
				var value2, _ = strconv.Atoi(each_ln2)
				var value3, _ = strconv.Atoi(each_ln3)

				if value1+value2+value3 == 2020 {
					fmt.Println(value1 * value2 * value3)
				}
			}
		}
	}
}
