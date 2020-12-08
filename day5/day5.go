package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day5_in")

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
	var max int64
	for _, each_ln1 := range text {
		fromBack := each_ln1[0:7]
		leftRight := each_ln1[8:10]

		binary := strings.Replace(fromBack, "F", "0", 7)
		binary = strings.Replace(binary, "B", "1", 7)
		numOfRow, _ := strconv.ParseInt(binary, 2, 64)

		binary = strings.Replace(leftRight, "L", "0", 3)
		binary = strings.Replace(binary, "R", "1", 3)
		numOfCol, _ := strconv.ParseInt(binary, 2, 64)

		boardingPass := numOfRow*8 + numOfCol
		if boardingPass > max {
			max = boardingPass
		}
	}
	fmt.Println(max)
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	var counters []int
	for _, each_ln1 := range text {
		fromBack := each_ln1[0:7]
		leftRight := each_ln1[8:10]

		binary := strings.Replace(fromBack, "F", "0", 7)
		binary = strings.Replace(binary, "B", "1", 7)
		numOfRow, _ := strconv.ParseInt(binary, 2, 8)

		binary = strings.Replace(leftRight, "L", "0", 3)
		binary = strings.Replace(binary, "R", "1", 3)
		numOfCol, _ := strconv.ParseInt(binary, 2, 8)

		boardingPass := int(numOfRow)*8 + int(numOfCol)
		counters = append(counters, boardingPass)
	}
	sort.Ints(counters)
	fmt.Println(getMissingNo(counters))

}

func getMissingNo(a []int) int {
	mySeat := 0
	for i, sn := range a {
		if i == len(a)-1 {
			continue
		}
		if a[i+1] != sn+1 {
			if sn+1 > 64*8 {
				mySeat = sn + 1
				break
			}
		}
	}
	return mySeat
}

func contains(a []int, n int) bool {
	for _, value := range a {
		if value == n {
			return true
		}
	}
	return false
}
