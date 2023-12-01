package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	monkey string
	items []int
	operation string
	arg_n int
	arg_s string
	test int
	ifTrue int
	ifFalse int
	inspections int
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func main(){

	monkeys := make([]Monkey, 0)
	monkey := ""
	items := make([]int, 0)
	item := 0
	operation := ""
	arg_n := 0
	arg_s := ""
	test := 0
	ifTrue := 0
	ifFalse := 0
	bigLimit := 1

	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	readFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
	fs := bufio.NewScanner(readFile)
	
	for fs.Scan() {
		line := strings.Fields(fs.Text())

		if len(line) > 0{
			if line[0] == "Monkey" {
				monkey = line[1][0:1]
			} else if line[0] == "Starting" {
				items = make([]int, 0)
				for _, i := range line[2:] {
					item, _ = strconv.Atoi(i[0:2])
					items = append([]int{item}, items...)
				}
			} else if line[0] == "Operation:" {
				if line[5] == "old" {
					arg_s = "old"
				} else {
					arg_s = ""
					arg_n, _ = strconv.Atoi(line[5])
				}
				operation = line[4]
			} else if line[0] == "Test:" {
				test, _ = strconv.Atoi(line[3])
				bigLimit *= test
			} else if line[1] == "true:" {
				ifTrue, _ = strconv.Atoi(line[5])
			} else if line[1] == "false:" {
				ifFalse, _ = strconv.Atoi(line[5])
			}
		} else {
			monkeys = append(monkeys, Monkey{monkey, items, operation, arg_n, arg_s, test, ifTrue, ifFalse, 0})
		}
	}
	monkeys = append(monkeys, Monkey{monkey, items, operation, arg_n, arg_s, test, ifTrue, ifFalse, 0})

	// Part 1
	// rounds := 20
	// Part 2
	rounds := 10000

	for i := 0; i < rounds; i++ {
		for indx, m := range monkeys {
			for j := len(m.items)-1; j >= 0; j-- {
				var WL int
				if m.arg_s == "old" {
					WL = ops[m.operation](m.items[j], m.items[j])
				} else {
					WL = ops[m.operation](m.items[j], m.arg_n)
				}

				// Part 1
				// WL = WL/3
				// Part 2
				WL = WL%bigLimit

				if WL % m.test == 0 {
					monkeys[m.ifTrue].items = append([]int{WL}, monkeys[m.ifTrue].items...)
					monkeys[indx].inspections += 1
				} else {
					monkeys[m.ifFalse].items = append([]int{WL}, monkeys[m.ifFalse].items...)
					monkeys[indx].inspections += 1
				}
				monkeys[indx].items = monkeys[indx].items[:len(monkeys[indx].items)-1]
			}
		}
	}

	max_1 := 0
	max_2 := 0
	for _, m := range monkeys {
		fmt.Println("Monkey", m.monkey, "inspected", m.inspections, "items.")
		if m.inspections > max_2 {
			max_2 = m.inspections
		} 
		if max_2 > max_1 {
			max_1, max_2 = max_2, max_1
		}
	}
	fmt.Println(max_1 * max_2)
}