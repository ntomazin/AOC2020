package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	file, err := os.Open("day16_in")

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
	counterInvalid := 0
	nearbyTickets := 0
	ranges := make(map[int]int)
	for i, each_ln1 := range text {
		if each_ln1 == "" {
			nearbyTickets = i + 4
			break
		}
		var lower1, upper1, lower2, upper2 int
		s := strings.Split(each_ln1, ":")
		if n, err := fmt.Sscanf(s[1], " %d-%d or %d-%d", &lower1, &upper1, &lower2, &upper2); n != 4 || err != nil {
			panic(fmt.Sprint(n, err))
		}

		for j := lower1; j <= upper1; j++ {
			ranges[j] = 1
		}
		for j := lower2; j <= upper2; j++ {
			ranges[j] = 1
		}

	}
	for _, each_ln1 := range text[nearbyTickets:] {
		numbers := strings.Split(each_ln1, ",")
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			if _, ok := ranges[num]; !ok {
				counterInvalid += num
			}
		}
	}
	fmt.Println(counterInvalid)

}

func part2(text []string) {
	input, _ := ioutil.ReadFile("day16_in")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	rules := map[string][]int{}
	for _, s := range strings.Split(split[0], "\n") {
		rule := strings.Split(s, ": ")
		rules[rule[0]] = make([]int, 4)
		fmt.Sscanf(rule[1], "%d-%d or %d-%d", &rules[rule[0]][0], &rules[rule[0]][1], &rules[rule[0]][2], &rules[rule[0]][3])
	}

	indices := map[string]map[int]struct{}{}
	for k := range rules {
		indices[k] = map[int]struct{}{}
		for i := 0; i < len(rules); i++ {
			indices[k][i] = struct{}{}
		}
	}

	part1 := 0
tickets:
	for _, s := range strings.Split(split[2], "\n")[1:] {
	fields:
		for _, s := range strings.Split(s, ",") {
			n, _ := strconv.Atoi(s)
			for _, v := range rules {
				if n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3] {
					continue fields
				}
			}
			part1 += n
			continue tickets
		}

		for i, s := range strings.Split(s, ",") {
			for k, v := range rules {
				if n, _ := strconv.Atoi(s); !(n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3]) {
					delete(indices[k], i)
				}
			}
		}
	}

	part2 := 1
	for len(indices) > 0 {
		for k, v := range indices {
			if len(v) != 1 {
				continue
			}

			for i := range v {
				for k := range indices {
					delete(indices[k], i)
				}
				delete(indices, k)

				if strings.HasPrefix(k, "departure") {
					n, _ := strconv.Atoi(strings.Split(strings.Split(split[1], "\n")[1], ",")[i])
					part2 *= n
				}
			}
		}
	}
	fmt.Println(part2)
}
