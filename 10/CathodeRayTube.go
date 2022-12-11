package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func incrementCycleCount(cyclesCount, X, sum *int) {
	*cyclesCount++
	if *cyclesCount == 20 || 
	   *cyclesCount == 60 ||
	   *cyclesCount == 100 || 
	   *cyclesCount == 140 || 
	   *cyclesCount == 180 || 
	   *cyclesCount == 220 {
		*sum = *sum + (*X)*(*(cyclesCount))
		//fmt.Println("Cycles:", *cyclesCount, "X:", *X, "strength:", *X*(*cyclesCount), "sum:", *sum)
	}
}

func main(){

	cyclesCount := 0
	X := 1
	sum := 0

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

	fs := bufio.NewScanner(readFile)

	for fs.Scan() {
		line := strings.Fields(fs.Text())
		incrementCycleCount(&cyclesCount, &X, &sum)
		if len(line) == 2 {
			addx, _ := strconv.Atoi(line[1])
			incrementCycleCount(&cyclesCount, &X, &sum)
			X += addx
		}

	}
	
	fmt.Println("Cycles:", cyclesCount, "X:", X, "Sum:", sum)

	readFile.Close()
}