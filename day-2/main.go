package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1(12, 13, 14)
	part2()
}

func part1(red int, green int, blue int) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		valid := true
		split := strings.Split(line, ":")
		game, err := strconv.Atoi(strings.Replace(split[0], "Game ", "", 1))
		if err != nil {
			log.Printf("failed to convert game string to a number: %s\n", err)
			continue
		}

		sets := strings.Split(split[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			r, g, b := getCounts(cubes)
			if r > red || g > green || b > blue {
				valid = false
				break
			}
		}

		if valid {
			total += game
		}
	}

	log.Printf("part one total: %d\n", total)
}

func part2() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, ":")
		_, err := strconv.Atoi(strings.Replace(split[0], "Game ", "", 1))
		if err != nil {
			log.Printf("failed to convert game string to a number: %s\n", err)
			continue
		}

		rMax := 0
		gMax := 0
		bMax := 0

		sets := strings.Split(split[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			r, g, b := getCounts(cubes)
			if r > rMax {
				rMax = r
			}

			if g > gMax {
				gMax = g
			}

			if b > bMax {
				bMax = b
			}
		}

		total += rMax * gMax * bMax
	}

	log.Printf("part two total: %d\n", total)
}

func getCounts(cubes []string) (int, int, int) {
	red := 0
	green := 0
	blue := 0

	for _, color := range cubes {
		if strings.Contains(color, "red") {
			r, err := strconv.Atoi(strings.Replace(strings.TrimSpace(color), " red", "", 1))
			if err != nil {
				log.Printf("failed to convert red cubes to number: %s\n", err)
			}

			red = r
		} else if strings.Contains(color, "green") {
			g, err := strconv.Atoi(strings.Replace(strings.TrimSpace(color), " green", "", 1))
			if err != nil {
				log.Printf("failed to convert green cubes to number: %s\n", err)
			}

			green = g
		} else if strings.Contains(color, "blue") {
			b, err := strconv.Atoi(strings.Replace(strings.TrimSpace(color), " blue", "", 1))
			if err != nil {
				log.Printf("failed to convert blue cubes to number: %s\n", err)
			}

			blue = b
		} else {
			log.Printf("unknown color provided: %s\n", color)
		}
	}

	return red, green, blue
}
