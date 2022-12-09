package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []rune

// Push adds an element to the stack
func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

// Pop removes the top element from the stack and returns it
func (s *Stack) Pop() rune {
	l := len(*s)
	if l == 0 {
		return 0
	}

	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() rune {
	l := len(*s)
	if l == 0 {
		return 0
	}

	return (*s)[l-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push in the beginning of the stack
func (s *Stack) PushFront(v rune) {
	if (s.IsEmpty()) {
		s.Push(v)
	} else {
		*s = append([]rune{v}, *s...)
	}
}

// PopSlice removes the top n elements from the stack and returns it
func (s *Stack) PopSlice(n int) []rune {
	l := len(*s)
	if l == 0 {
		return nil
	}

	if n > l {
		n = l
	}

	v := (*s)[l-n:]
	*s = (*s)[:l-n]
	return v
}

// PushSlice adds an slice to the stack
func (s *Stack) PushSlice(v []rune) {
	*s = append(*s, v...)
}

func main() {
	
	stacks_one := make([]Stack, 10)
	stacks_two := make([]Stack, 10)

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

    fs := bufio.NewScanner(readFile)
 
    fs.Split(bufio.ScanLines)

	for fs.Scan() {
		if fs.Text() != "" {
			for i, r := range fs.Text() {
				if r >= 60 && r <= 90 {
					stacks_one[i/4].PushFront(r)
				} 
			}
		} else {
			break
		}
	}
	
	for i := range stacks_one {
		stacks_two[i] = make([]rune, len(stacks_one[i]))
		copy(stacks_two[i], stacks_one[i])
	}

	for fs.Scan() {
		var cargoCount, fromCrate, toCrate int
		fmt.Sscanf(fs.Text(), "move %d from %d to %d", &cargoCount, &fromCrate, &toCrate)
		
		for count := 0; count < cargoCount; count++ {
			stacks_one[toCrate-1].Push(stacks_one[fromCrate-1].Pop())
		} 

		stacks_two[toCrate-1].PushSlice(stacks_two[fromCrate-1].PopSlice(cargoCount))
		//fmt.Println(stacks_two)
	}
	readFile.Close()

	for _, s := range stacks_one{
		fmt.Print(string(s.Peek()))
	}
	fmt.Print("\n")

	for _, s := range stacks_two{
		fmt.Print(string(s.Peek()))
	}
	fmt.Print("\n")
}