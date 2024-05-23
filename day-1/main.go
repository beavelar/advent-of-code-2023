package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers map[string]int

func main() {
	numbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

    part1()
    part2(numbers)
}

func part1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err)
	}

	content := string(data)
	split := strings.Split(content, "\n")

	total := 0
	for _, line := range split {
		left := 0
		right := 0
        reset := true

		for _, rChar := range line {
			char := string(rChar)
			num, err := strconv.Atoi(char)
			if err != nil {
                continue
			}

			if reset {
				left = num
				reset = false
			}

			right = num
		}

		if reset {
			continue
		}

		reset = true
		cmb := fmt.Sprintf("%d%d", left, right)
		num, err := strconv.Atoi(cmb)
		if err != nil {
			log.Printf("failed to combine digits: %s\n", err)
			continue
		}

		total += num
		continue
	}

	log.Printf("part one total: %d\n", total)
}

func part2(numbers map[string]int) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err)
	}

	content := string(data)
	split := strings.Split(content, "\n")

	total := 0
	for _, line := range split {
		left := 0
		right := 0
        reset := true

		for idx, rChar := range line {
			char := string(rChar)
			num, err := strconv.Atoi(char)
			if err != nil {
                dig, valid := strToInt(line[idx:], numbers)
                if valid {
                    num = dig
                } else {
                    continue
                }
			}

			if reset {
				left = num
				reset = false
			}

			right = num
		}

		if reset {
			continue
		}

		reset = true
		cmb := fmt.Sprintf("%d%d", left, right)
		num, err := strconv.Atoi(cmb)
		if err != nil {
			log.Printf("failed to combine digits: %s\n", err)
			continue
		}

		total += num
		continue
	}

	log.Printf("part two total: %d\n", total)
}

func strToInt(line string, numbers map[string]int) (int, bool) {
	for key, value := range numbers {
		if len(line) >= len(key) && line[:len(key)] == key {
			return value, true
		}
	}

	return 0, false
}
