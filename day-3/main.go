package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	NUM = iota
	SYM
	NA
)

func main() {
	part1()
}

func part1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")
	lLen := len(lines[0])

	cStart := -1
	total := 0
	for lidx, line := range lines {
		for cidx, r := range line {
			cType := getType(r)
			if cStart == -1 && cType == NUM {
				if cidx == 0 {
					cStart = cidx
				} else {
					cStart = cidx - 1
				}
			} else if cStart != -1 && (cType != NUM || cidx == lLen-1) {
				var tLines []string
				if lidx == 0 {
					tLines = append(tLines, lines[0][cStart:cidx], lines[1][cStart:cidx])
				} else if lidx == len(lines)-2 {
					tLines = append(tLines, lines[lidx][cStart:cidx], lines[lidx-1][cStart:cidx])
				} else {
					tLines = append(tLines, lines[lidx-1][cStart:cidx], lines[lidx][cStart:cidx], lines[lidx+1][cStart:cidx])
				}

				num, valid := getNumber(tLines)
				if valid {
					total += num
				}

				cStart = -1
			}
		}
	}

	log.Printf("part one total: %d\n", total)
}

func getNumber(slice []string) (int, bool) {
	valid := false
	lLen := len(slice[0])
	numS := ""

	bChars := ""
	var dLine string
	if len(slice) == 2 {
		dLine = slice[0]
		bChars += slice[1] + slice[0][0:1] + slice[0][lLen-1:lLen]
	} else {
		dLine = slice[1]
		bChars += slice[0] + slice[2] + slice[1][0:1] + slice[1][lLen-1:lLen]
	}

	for _, char := range bChars {
		if getType(char) == SYM {
			valid = true
			break
		}
	}

	for _, char := range dLine {
		if getType(char) == NUM {
			numS += string(char)
		}
	}

	num, err := strconv.Atoi(numS)
	if err != nil {
		log.Printf("failed to convert number %s to int: %s", numS, err)
		return 0, false
	}

	log.Printf("number retrieved: %d, validity: %t", num, valid)
	return num, valid
}

func getType(char rune) int {
	if unicode.IsDigit(char) {
		return NUM
	}

	if string(char) == "." {
		return NA
	}

	return SYM
}
