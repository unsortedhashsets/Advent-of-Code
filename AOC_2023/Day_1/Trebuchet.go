package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValue_Two(line string) int {
	var result []int
	currentWord := ""
	output := 0
	for _, r := range line {
		if unicode.IsDigit(r) {
			result = append(result, int(r-'0'))
		} else {
			currentWord += string(r)
			test := false
			for word, digit := range digitMap {
				if strings.HasPrefix(word, currentWord) {
					test = true
					if word == currentWord {
						result = append(result, digit)
						currentWord = string(r)
					}
					break
				}
			}
			if !test {
				currentWord = currentWord[1:]
			}
		}
	}
	if len(result) > 0 {
		output = result[0]*10 + result[len(result)-1]
	}
	return output
}

func getCalibrationValue_One(line string) int {
	var result []int
	output := 0
	for _, r := range line {
		if unicode.IsDigit(r) {
			result = append(result, int(r-'0'))
		}
	}
	if len(result) > 0 {
		output = result[0]*10 + result[len(result)-1]
	}
	return output
}

func processLines(scanner *bufio.Scanner) (int, int) {
	one, two := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		one += getCalibrationValue_One(line)
		two += getCalibrationValue_Two(line)
	}
	return one, two
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	one, two := processLines(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part One: %d\n", one)
	fmt.Printf("Part Two: %d\n", two)
}
