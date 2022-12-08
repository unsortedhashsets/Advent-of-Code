package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"sort"
)

func insertInSortedArray(caloriesArray *[]int, newCalories *int) {
	if len(*caloriesArray) == 0 {
		*caloriesArray = append(*caloriesArray, *newCalories)
	} else {
		i := sort.Search(len(*caloriesArray), func(i int) bool { return (*caloriesArray)[i] >= *newCalories })
		*caloriesArray = append(*caloriesArray, 0)
		copy((*caloriesArray)[i+1:], (*caloriesArray)[i:])
		(*caloriesArray)[i] = *newCalories
	}
	
}

func main() {
	
	caloriesArray := []int{}
	maxCalories := 0
	newCalories := 0

    readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			insertInSortedArray(&caloriesArray, &newCalories)
			newCalories = 0
		} else {
			inputValue, e := strconv.Atoi(fileScanner.Text())
			if e != nil {
				fmt.Println(e)
			} else {
				newCalories += inputValue
				if newCalories > maxCalories {
					maxCalories = newCalories
				}
			}
		}      
    }

    readFile.Close()
	
	lastThree := caloriesArray[len(caloriesArray) - 1] + caloriesArray[len(caloriesArray) - 2] + caloriesArray[len(caloriesArray) - 3]
	
	fmt.Println(maxCalories)
	fmt.Println(lastThree)
}