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
	file, err := os.Open("day4_in")

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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func part1(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields_counter := 0
	for _, each_ln1 := range text {
		if string(each_ln1) == "" {
			if fields_counter == len(fields) {
				counter++
			}
			fields_counter = 0
			continue
		}

		s := strings.Split(each_ln1, " ")
		for _, pair := range s {
			key := strings.Split(pair, ":")[0]

			//value := strings.Split(pair, ":")[1]
			if contains(fields, key) {
				fields_counter++
			}
		}
	}
	if fields_counter == len(fields) {
		counter++
	}

	fmt.Println(counter)
}

func part2(text []string) {
	// and then a loop iterates through
	// and prints each of the slice values.
	counter := 0
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields_counter := 0
	for _, each_ln1 := range text {
		if string(each_ln1) == "" {
			if fields_counter == len(fields) {
				counter++
			}
			fields_counter = 0
			continue
		}

		s := strings.Split(each_ln1, " ")
		for _, pair := range s {
			key := strings.Split(pair, ":")[0]
			value := strings.Split(pair, ":")[1]
			if contains(fields, key) {
				if testValid(key, value) {
					fields_counter++
				}
			}
		}
	}
	if fields_counter == len(fields) {
		counter++
	}

	fmt.Println(counter)
}

func testValid(key string, value string) bool {
	if key == "byr" {
		i, _ := strconv.Atoi(value)
		if i <= 2002 && i >= 1920 {
			return true
		}
	}
	if key == "iyr" {
		i, _ := strconv.Atoi(value)
		if i <= 2020 && i >= 2010 {
			return true
		}
	}
	if key == "eyr" {
		i, _ := strconv.Atoi(value)
		if i <= 2030 && i >= 2020 {
			return true
		}
	}
	if key == "hgh" {
		if strings.HasSuffix(value, "cm") {
			t := strings.Replace(value, "cm", "", 1)
			i, err := strconv.Atoi(t)
			if err != nil {
				return false
			}
			if 150 <= i && i <= 193 {
				return true
			}
		}
		if strings.HasSuffix(value, "in") {
			t := strings.Replace(value, "in", "", 1)
			i, err := strconv.Atoi(t)
			if err != nil {
				return false
			}
			if 59 <= i && i <= 76 {
				return true
			}
		}
	}
	if key == "hcl" {
		matched, _ := regexp.MatchString(`^#[a-f0-9]{6}$`, value)
		return matched
	}
	if key == "ecl" {
		possibilities := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		return contains(possibilities, value)
	}
	if key == "pid" {
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if len(value) == 9 {
			return true
		}
	}

	return false
}
