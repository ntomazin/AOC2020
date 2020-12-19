package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("day14_in")

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
	// and then a loop iterates through
	// and prints each of the slice values.
	memory := make(map[int]uint64)
	var mask string
	for _, each_ln1 := range text {
		s := strings.Split(each_ln1, "=")
		if s[0] == "mask " {
			mask = s[1][1:len(s[1])]
		} else {
			memIdx := 0
			var toSet uint64 = 0
			if n, err := fmt.Sscanf(each_ln1, "mem[%d] = %d", &memIdx, &toSet); n != 2 || err != nil {
				panic(fmt.Sprint(n, err))
			}

			memory[memIdx] = applyMask(mask, toSet)
		}
	}
	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	fmt.Println(total)
}

func part2(text []string) {
	memory := make(map[uint64]uint64)
	mask := ""
	masks := []maskWithFloats{}
	for _, each_ln1 := range text {
		s := strings.Split(each_ln1, "=")
		if s[0] == "mask " {
			mask = s[1][1:len(s[1])]
			masks = maskPermut(mask)

		} else {
			var memIdx uint64 = 0
			var toSet uint64 = 0
			if n, err := fmt.Sscanf(each_ln1, "mem[%d] = %d", &memIdx, &toSet); n != 2 || err != nil {
				panic(fmt.Sprint(n, err))
			}

			for _, maska := range masks {
				target := applyMaskP2(maska, memIdx)
				memory[target] = toSet
			}
		}
	}
	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	fmt.Println(total)
}

func applyMask(mask string, input uint64) uint64 {
	for i, v := range mask {
		switch v {
		case 'X':
		case '0':

			input &= ^(1 << (35 - i))
		case '1':
			input |= (1 << (35 - i))
		}
	}
	return input
}

func maskPermut(mask string) []maskWithFloats {
	mask = strings.TrimPrefix(mask, "mask = ")
	var out []maskWithFloats
	num := strings.Count(mask, "X")
	var realNum uint64 = 0xFFFFFFFFFFFFFFFF >> (64 - num)
	printFMask := fmt.Sprintf("%%0%db", num)
	nextMask := strings.Builder{}
	nextMask.Grow(len(mask))
	for i := 0; uint64(i) <= realNum; i++ {
		bitCount := 0
		bits := fmt.Sprintf(printFMask, i)
		floats := [][2]int{}
		for i, chr := range mask {
			switch chr {
			case 'X':
				nextMask.WriteByte(bits[bitCount])
				floats = append(floats, [2]int{35 - i, int(bits[bitCount]) - 48})
				bitCount++
			default:
				nextMask.WriteRune(chr)
			}
		}

		out = append(out, maskWithFloats{mask: nextMask.String(), floats: floats})
		nextMask.Reset()
	}

	return out
}

func setBit(num uint64, target, bit int) uint64 {
	switch bit {
	case 0:
		num &= ^(1 << target)
	case 1:
		num |= 1 << target
	default:
		panic("Tried to set a bit that is neither 0 nor 1")
	}
	return num
}

func applyMaskP2(mask maskWithFloats, input uint64) uint64 {
	mask.mask = strings.TrimPrefix(mask.mask, "mask = ")
	for i, v := range mask.mask {
		switch v {
		case '1':
			input |= (1 << (35 - i))
		}
		for _, v := range mask.floats {
			if v[0] == i {
				input = setBit(input, v[0], v[1])
			}
		}
	}
	return input
}

type maskWithFloats struct {
	mask   string
	floats [][2]int
}
