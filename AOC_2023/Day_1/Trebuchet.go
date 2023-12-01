package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func replaceWordsWithDigits(line string) string {
	result := ""
	currentWord := ""
	for _, r := range line {
		if unicode.IsDigit(r) {
			result += string(r)
		} else {
			currentWord += string(r)
			test := false
			for word, digit := range digitMap {
				if strings.HasPrefix(word, currentWord) {
					test = true
					if word == currentWord {
						result += digit
						currentWord = string(r)
					}
					break
				}
			}
			if !test {
				result += string(currentWord[0])
				currentWord = currentWord[1:]
			}
		}
	}
	return result
}

func getCalibrationValue(line string) (int, error) {
	firstDigit, lastDigit := -1, -1
	for _, r := range line {
		if unicode.IsDigit(r) {
			if firstDigit == -1 {
				firstDigit = int(r - '0')
			}
			lastDigit = int(r - '0')
		}
	}
	if firstDigit != -1 && lastDigit != -1 {
		return strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
	}
	return 0, fmt.Errorf("no calibration value found")
}

func processLines(scanner *bufio.Scanner) (int, int) {
	one := 0
	two := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibrationValueOne, err := getCalibrationValue(line)
		line = replaceWordsWithDigits(line)
		calibrationValueTwo, err := getCalibrationValue(line)
		if err == nil {
			one += calibrationValueOne
			two += calibrationValueTwo
		}
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
