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

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day9_in")

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
	var queue []int
	isValid := false
	for _, each_ln1 := range text {
		num, _ := strconv.Atoi(each_ln1)

		if len(queue) < 25 {
			queue = append(queue, num)
			continue
		}
		if len(queue) >= 25 {
			isValid = false
			for _, elem1 := range queue {
				if isValid {
					break
				}
				for _, elem2 := range queue {
					if elem1+elem2 == num {
						queue = queue[1:len(queue)] // Dequeue
						queue = append(queue, num)
						isValid = true
					}
				}
			}

		}
		if !isValid {
			fmt.Println(each_ln1)
			return
		}
	}
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	xmas := make([]int, len(text))
	for i, s := range text {
		xmas[i], _ = strconv.Atoi(s)
	}
	invalidNumber := 15690279
	for i := 0; i < len(xmas); i++ {
		for j := i + 1; j < len(xmas); j++ {
			sum := 0
			for _, v := range xmas[i : j+1] {
				sum += v
			}
			if sum == invalidNumber {
				sort.Ints(xmas[i : j+1])
				fmt.Println(xmas[i] + xmas[j])
				return
			}
		}
	}
}
