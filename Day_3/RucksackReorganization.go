package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	
    scoreOne := 0

    
    lineOne := ""
    lineTwo := ""
    lineThree := ""
    scoreTwo := 0

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    

    for fileScanner.Scan() {

        lineLength := len(fileScanner.Text())
        compartmentOne := fileScanner.Text()[: lineLength / 2]
        compartmentTwo := fileScanner.Text()[lineLength / 2 : lineLength]
        alreadyWasOne := ""

        for _, s := range compartmentOne {
            // Part One
            if (strings.Contains(compartmentTwo, string(s))){
                if !(strings.Contains(alreadyWasOne, string(s))){
                    alreadyWasOne += string(s)

                    if (s <= 91) {
                        scoreOne += int(s) - 38
                    } else {
                        scoreOne += int(s) - 96
                    }
                }
            }
        }
        // Part Two
        if (lineThree == "") {
            if (lineTwo == "") {
                if (lineOne == "") {
                    lineOne = fileScanner.Text()
                } else {
                    lineTwo = fileScanner.Text()
                }
            } else {
                lineThree = fileScanner.Text()
                alreadyWasTwo := ""
                for _, s := range lineThree {
                    if (strings.Contains(lineTwo, string(s))){
                        if (strings.Contains(lineOne, string(s))){
                            if !(strings.Contains(alreadyWasTwo, string(s))){
                                alreadyWasTwo += string(s)
                                if (s <= 91) {
                                    scoreTwo += int(s) - 38
                                } else {
                                    scoreTwo += int(s) - 96
                                }
                            }
                        }
                    }
                }
                lineOne = ""
                lineTwo = ""
                lineThree = ""
            }
        }
    }

    readFile.Close()

	fmt.Println(scoreOne)
    fmt.Println(scoreTwo)
}