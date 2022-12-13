package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func scanInput() [][]any  {
	readFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
	fs := bufio.NewScanner(readFile)

	packets := [][]any{}
	pair := []any{}

	for fs.Scan() {
		if fs.Text() == "" {
			packets = append(packets, pair)
			pair = []any{}
		} else {
			var p any
			err := json.Unmarshal([]byte(fs.Text()), &p)
			if err != nil {
				log.Fatal(err)
			}
			pair = append(pair, p)
		}
	}

	return append(packets, pair)
}

func DistressSignalPartOne(packets [][]any) int {
	result := 0
	for i, pair := range packets {
		if compare(pair[0], pair[1]) <= 0 {
			result += i + 1
		}
	}
	return result
}

func compare(left any, right any) int {
	_, valueLeft := left.(float64)
	_, valueRight := right.(float64)
	
	if valueLeft && valueRight {
		return int(left.(float64)) - int(right.(float64))
	}

	if valueLeft {
		left = []any{left}
	}
	if valueRight {
		right = []any{right}
	}

	if len(left.([]any)) == 0 || len(right.([]any)) == 0 {
		return len(left.([]any)) - len(right.([]any))
	}

	result := compare(left.([]any)[0], right.([]any)[0])
	
	if result == 0 {
		left_next := left.([]any)[1:]
		right_next := right.([]any)[1:]

		if len(left_next) == 0 || len(right_next) == 0 {
			return len(left_next) - len(right_next)
		}
		return compare(left_next, right_next)
	}

	return result
}

func DistressSignalPartTwo(packets [][]any) int {
	new := []any{}
	for _, pair := range packets {
		new = append(new, pair...)
	}

	var divider1 any
	err := json.Unmarshal([]byte("[[2]]"), &divider1)
	if err != nil {
		log.Fatal(err)
	}

	var divider2 any
	err = json.Unmarshal([]byte("[[6]]"), &divider2)
	if err != nil {
		log.Fatal(err)
	}

	new = append(new, []any{divider1, divider2}...)
	sort.Slice(new, func(i, j int) bool {
		return compare(new[i], new[j]) <= 0
	})

	result := 1
	for i, packet := range new {
		packet, err := json.Marshal(packet)
		if err != nil {
			log.Fatal(err)
		}
		packetString := string(packet)
		if packetString == "[[2]]" || packetString == "[[6]]" {
			result *= i + 1
		}
	}

	return result
}

func main() {
	packets := scanInput()
	fmt.Println("Solve Part 1:", DistressSignalPartOne(packets))
	fmt.Println("Solve Part 2:", DistressSignalPartTwo(packets))
}