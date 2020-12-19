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
	file, err := os.Open("day19_in")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	startTime := time.Now()
	part1(text)
	fmt.Println("Part 1 took:", time.Since(startTime))

	startTime = time.Now()
	//part2(text)
	fmt.Println("Part 2 took:", time.Since(startTime))
}

func part1(text []string) {
	rules := map[string]string{}
	part2 := 0
	for i, line := range text {
		if line == "" {
			part2 = i
			break
		}
		s := strings.Split(line, " ")
		key := strings.Split(s[0], ":")[0]
		tempString := "( "
		for _, v := range s[1:] {
			if v != "|" {
				tempString += v + " "
			} else {
				tempString += ") | ( "
			}
		}
		tempString += ")"
		rules[key] = tempString
	}
	fmt.Println(rules)
	for {
		newRules := make(map[string]string)
		hasChanged := false
		for k, v := range rules {
			replacementString := ""
			for _, v3 := range strings.Split(v, " ") {
				if v3 != "\"a\"" && v3 != "\"b\"" &&
					v3 != "(" && v3 != ")" &&
					v3 != "|" && v3 != "" {
					replacementString += rules[v3]
					hasChanged = true
				}
				if v3 == "\"a\"" {
					replacementString += " \"a\" "
				}
				if v3 == "\"b\"" {
					replacementString += " \"b\" "
				}
				if v3 == "|" {
					replacementString += " | "
				}
			}
			newRules[k] = replacementString
			//fmt.Println(newRules)
		}
		rules = newRules
		if !hasChanged {
			break
		}
	}
	fmt.Println(rules)
	for i, line := range text[part2:] {
		if string(i) == string(line[0]) {
			continue
		}
	}

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

		if i == 30000000-1 {
			fmt.Println(last)
		}
	}
}
