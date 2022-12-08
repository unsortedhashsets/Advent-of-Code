package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	
	scoreOne := 0
	scoreTwo := 0

	readFile, err := os.Open("input.txt")
	mappingOne := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	mappingTwo := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}
	
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
		scoreOne += mappingOne[fileScanner.Text()]
		scoreTwo += mappingTwo[fileScanner.Text()]
    }

    readFile.Close()

	fmt.Println(scoreOne)
	fmt.Println(scoreTwo)
}