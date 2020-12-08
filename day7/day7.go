package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day7_in")

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
	//part2(text)
}

type Bag struct {
	name     string
	quantity int
}

func part1(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	allBags := make(map[string]Bag)
	re := regexp.MustCompile(`(\d+) (.+?) bag`)

	for _, each_ln1 := range text {
		s := strings.Split(each_ln1, "contain")
		mainBag := strings.Join(strings.Split(s[0], " ")[0:1], " ")
		otherBags := strings.Split(s[1], ",")

		for _, bag := range otherBags {
			fmt.Println(bag)

			matches := re.FindStringSubmatch(bag)
			fmt.Println(matches[0])
			fmt.Println(matches[1])

			quantity, _ := strconv.Atoi(matches[0])
			allBags[mainBag] = Bag{
				name:     matches[1],
				quantity: quantity,
			}
		}
	}
	fmt.Println(counter)
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	for _, each_ln1 := range text {
		s := strings.Split(each_ln1, " ")
		num := strings.Split(s[0], "-")
		letter := strings.Split(s[1], ":")[0]
		password := s[2]
		num1, _ := strconv.Atoi(num[0])
		num2, _ := strconv.Atoi(num[1])
		if string(password[num1-1]) == letter || string(password[num2-1]) == letter {
			counter++
		}

		if string(password[num1-1]) == letter && string(password[num2-1]) == letter {
			counter--
		}

	}
	fmt.Println(counter)
}
