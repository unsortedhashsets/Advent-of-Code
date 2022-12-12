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

func ScanMap(Map *[][]rune, Start *Coordinate, End *Coordinate){
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
				s = 96
			} else if s == 'E'{
				*End = Coordinate{i, len(*Map)}
				s = 123
			}
			row = append(row, s)
		}
		*Map = append(*Map, row)
	}
	readFile.Close()
}

func FindWayPartOne(Map [][]rune, Start Coordinate, End Coordinate){
	Visited := make(map[Coordinate]bool)
	PQueue := []Coordinate{Start}
	Steps := map[Coordinate]int{Start:0}

	for {	
		Point := PQueue[0]
		Visited[Point] = true
		PQueue = PQueue[1:]

		if Point == End{
			fmt.Println("Part One - Steps:", Steps[End])
			break
		}
		
		for _, xy := range [][]int{{1,0},{0,-1},{-1,0},{0,1}}{
			x, y := xy[0], xy[1]
			NewPoint := Coordinate{Point.x+x, Point.y+y}
			// Check if visited
			if !Visited[NewPoint] {
				// Check if in bounds
				if  NewPoint.x>=0 && 
					NewPoint.y>=0 &&
					NewPoint.x<len(Map[0]) &&
					NewPoint.y<len(Map) {
					// Check if h is within 1
					if Map[NewPoint.y][NewPoint.x]-Map[Point.y][Point.x]<=1{
						// Check if in queue
						if Steps[NewPoint] == 0{
							PQueue = append(PQueue, NewPoint)
							Steps[NewPoint] = Steps[Point]+1
						}
						if Steps[NewPoint] >= Steps[Point]+1{
							Steps[NewPoint] = Steps[Point]+1
						}
					}
				}
			}
		}
		sort.Slice(PQueue, func(i,j int)bool{
			return Steps[PQueue[i]] < Steps[PQueue[j]]
		})
	}
}

func main(){

	Map := make([][]rune, 0)
	var Start, End Coordinate

	ScanMap(&Map, &Start, &End)
	FindWayPartOne(Map, Start, End)

}