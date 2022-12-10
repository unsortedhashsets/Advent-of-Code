package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

func calculateDistance(forest [][]int, height int, i int, j int) int {
	return calculateDistanceLeft(forest, height, i, j) *
		   calculateDistanceRight(forest, height, i, j) *
		   calculateDistanceTop(forest, height, i, j) *
		   calculateDistanceBottom(forest, height, i, j)
}

func calculateDistanceLeft(forest [][]int, height int, i int, j int) int {
	distance := 0
	for k := j-1; k >= 0; k-- {
		if forest[i][k] >= height {
			distance++
			return distance
		}
		distance++
	}
	return distance
}

func calculateDistanceRight(forest [][]int, height int, i int, j int) int {
	distance := 0
	for k := j + 1; k < len(forest[i]); k++ {
		if forest[i][k] >= height {
			distance++
			return distance
		}
		distance++
	}
	return distance
}

func calculateDistanceTop(forest [][]int, height int, i int, j int) int {
	distance := 0
	for k := i - 1; k >= 0; k-- {
		if forest[k][j] >= height {
			distance++
			return distance
		}
		distance++
	}
	return distance
}

func calculateDistanceBottom(forest [][]int, height int, i int, j int) int {
	distance := 0
	for k := i + 1; k < len(forest); k++ {
		if forest[k][j] >= height {
			distance++
			return distance
		}
		distance++
	}
	return distance
}

func main(){

	forest := make([][]int, 0)

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

	fs := bufio.NewScanner(readFile)
	
	for fs.Scan() {
		row := make([]int, 0)
		for _, t := range fs.Text() {
			t_int, _ := strconv.Atoi(string(t))
			row = append(row, t_int)
		}
		forest = append(forest, row)
	}
	
	readFile.Close()

	visibleTrees := make(map[Coordinate]bool)

	highestTop := 0
	highestBottom := 0
	highestLeft := 0
	highestRight := 0
	for i :=0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			// visible on the left edge (highestLeft and right are edges)
			if j == 0{
				highestLeft = forest[i][j]
				highestRight = forest[i][len(forest[j])-1]
				visibleTrees[Coordinate{i, j}] = true
				visibleTrees[Coordinate{i, len(forest[j])-1}] = true
			} else {
				// left -> right
				if forest[i][j] > highestLeft {
					//fmt.Println("left -> right;", i, j, highestLeft, forest[i][j])
					highestLeft = forest[i][j]
					visibleTrees[Coordinate{i, j}] = true
				}
				// right -> left
				if forest[i][len(forest[i])-j-1] > highestRight {
					//fmt.Println("right -> left;", i, len(forest[i])-j-1, highestRight, forest[i][len(forest[i])-j-1])
					highestRight = forest[i][len(forest[i])-j-1]
					visibleTrees[Coordinate{i, len(forest[i])-j-1}] = true
				}
			}
			// visible on the top edge (highestTop and bottom are edges) (change axis)
			if j == 0 {
				highestTop = forest[j][i]
				highestBottom = forest[len(forest)-j-1][i]
				visibleTrees[Coordinate{j, i}] = true
				visibleTrees[Coordinate{len(forest[j])-1, i}] = true
			} else {
				// top -> bottom
				if forest[j][i] > highestTop {
					//fmt.Println("top -> bottom;", j, i, highestTop, forest[j][i])
					highestTop = forest[j][i]
					visibleTrees[Coordinate{j, i}] = true
				}
				// bottom -> top
				if forest[len(forest)-j-1][i] > highestBottom {
					//fmt.Println("bottom -> top;", len(forest)-j-1, i, highestTop, forest[len(forest)-i-1][j])
					highestBottom = forest[len(forest)-j-1][i]
					visibleTrees[Coordinate{len(forest)-j-1, i}] = true
				}
			}
		}
	}

	fmt.Println(len(visibleTrees))

	// Part Two

	score := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			tmp_score := calculateDistance(forest, forest[i][j], i, j)
			if tmp_score > score {
				score = tmp_score
			}
		}
	}

	fmt.Println(score)
}