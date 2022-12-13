package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {
	
	score_one := 0
	score_two := 0
	
	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
		pairs := strings.Split(fileScanner.Text(), ",")
		elve_one := strings.Split(pairs[0], "-")
		elve_two := strings.Split(pairs[1], "-")
		
		elve_one_start, _ := strconv.Atoi(elve_one[0])
		elve_one_end, _ := strconv.Atoi(elve_one[1])
		elve_two_start, _ := strconv.Atoi(elve_two[0])
		elve_two_end, _ := strconv.Atoi(elve_two[1])

		if elve_one_start >= elve_two_start && elve_one_end <= elve_two_end || elve_two_start >= elve_one_start && elve_two_end <= elve_one_end {	
			score_one++
		}

		if elve_one_start <= elve_two_end && elve_two_start <= elve_one_start || elve_two_start <= elve_one_end && elve_one_start <= elve_two_start {	
			score_two++
		}
			
    }

    readFile.Close()

	fmt.Println(score_one)
	fmt.Println(score_two)
}