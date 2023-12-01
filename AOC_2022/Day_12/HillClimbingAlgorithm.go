package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coordinate struct{
	x,y int
}

func ScanMap(Map *[][]rune, Start *Coordinate, Starts *[]Coordinate, End *Coordinate){
	readFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
	fs := bufio.NewScanner(readFile)
	
	for fs.Scan(){
		var row []rune
		for i, s := range fs.Text(){
			if s == 'S'{
				*Start = Coordinate{i, len(*Map)}
				s = 97
			} else if s == 'E'{
				*End = Coordinate{i, len(*Map)}
				s = 123
			} else if s == 'a'{
				*Starts = append(*Starts, Coordinate{i, len(*Map)})
			}
			row = append(row, s)
		}
		*Map = append(*Map, row)
	}
	readFile.Close()
}

func FindWayPartOne(Map *[][]rune, Start *Coordinate, End *Coordinate, Visited *map[Coordinate]bool, PQueue *[]Coordinate, Steps *map[Coordinate]int, min int) int {
	for {	
		if len(*PQueue) == 0{
			return min
		}
		Point := (*PQueue)[0]

		if min != 0 && (*Steps)[Point] > min{
			return min
		}

		*PQueue = (*PQueue)[1:]
		(*Visited)[Point] = true
		


		if Point == *End{
			return (*Steps)[*End]
		}
		
		for _, xy := range [][]int{{1,0},{0,-1},{-1,0},{0,1}}{
			x, y := xy[0], xy[1]
			NewPoint := Coordinate{Point.x+x, Point.y+y}
			// Check if visited
			if !(*Visited)[NewPoint] {
				// Check if in bounds
				if  NewPoint.x>=0 && 
					NewPoint.y>=0 &&
					NewPoint.x<len((*Map)[0]) &&
					NewPoint.y<len(*Map) {
					// Check if h is within 1
					if (*Map)[NewPoint.y][NewPoint.x]-(*Map)[Point.y][Point.x]<=1{
						// Check if in queue
						if (*Steps)[NewPoint] == 0{
							*PQueue = append(*PQueue, NewPoint)
							(*Steps)[NewPoint] = (*Steps)[Point]+1
						}
						if (*Steps)[NewPoint] >= (*Steps)[Point]+1{
							(*Steps)[NewPoint] = (*Steps)[Point]+1
						}
					}
				}
			}
		}
		sort.Slice(*PQueue, func(i,j int)bool{
			return (*Steps)[(*PQueue)[i]] < (*Steps)[(*PQueue)[j]]
		})
	}
}

func FindWayPartTwo(Map *[][]rune, Starts *[]Coordinate, End *Coordinate, Visited *map[Coordinate]bool, PQueue *[]Coordinate, Steps *map[Coordinate]int, min int) int {
	for _, Start := range *Starts {
		*Visited = make(map[Coordinate]bool)
		*PQueue = []Coordinate{Start}
		*Steps = map[Coordinate]int{Start:0}
		min = FindWayPartOne(Map, &Start, End, Visited, PQueue, Steps, min)
	}
	return min
}

func main(){

	Map := make([][]rune, 0)
	var Start, End Coordinate
	var Starts []Coordinate

	ScanMap(&Map, &Start, &Starts, &End)
	Visited := make(map[Coordinate]bool)
	PQueue := []Coordinate{Start}
	Steps := map[Coordinate]int{Start:0}

	pOne := FindWayPartOne(&Map, &Start, &End, &Visited, &PQueue, &Steps, 0)
	fmt.Println("Part One - Steps:", pOne)

	pTwo := FindWayPartTwo(&Map, &Starts, &End, &Visited, &PQueue, &Steps, pOne)
	fmt.Println("Part Two - Steps:", pTwo)
	
}