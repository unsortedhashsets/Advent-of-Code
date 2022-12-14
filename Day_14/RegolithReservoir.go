package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	// Use to update picture every 100ms (for small example)
	/*
	"os/exec"
	"time"
	*/
	
)

type Coordinate struct{
	x,y int
}

func scanInput(Cave *map[Coordinate]rune, min_X, max_X, max_Y *int) {
	readFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
	fs := bufio.NewScanner(readFile)

	for fs.Scan() {
		Coordinates := strings.Split(fs.Text(), " -> ")
		for i := range Coordinates[:len(Coordinates)-1]{
			from := strings.Split(Coordinates[i], ",")
			to := strings.Split(Coordinates[i+1], ",")
			from_X, _ := strconv.Atoi(from[0])
			from_Y, _ := strconv.Atoi(from[1]) 
			to_X, _ := strconv.Atoi(to[0])
			to_Y, _ := strconv.Atoi(to[1]) 

			(*Cave)[Coordinate{to_X, to_Y}] = '#'
			(*Cave)[Coordinate{from_X, from_Y}] = '#'

			if from_Y>*max_Y{
				*max_Y = from_Y
			}
			if from_X<*min_X{
				*min_X = from_X
			}
			if to_Y>*max_Y{
				*max_Y = to_Y
			}

			if to_X>*max_X{
				*max_X = to_X
			}	
			if to_X<*min_X{
				*min_X = to_X
			}

			for from_X != to_X || from_Y != to_Y{
				(*Cave)[Coordinate{from_X, from_Y}] = '#'
				switch {
				case from_X < to_X:
					from_X++
				case from_Y < to_Y :
					from_Y++
				case from_X > to_X:
					from_X--
				case from_Y > to_Y:
					from_Y--
				}
			}
		}
	}
}

func updateInput(Cave *map[Coordinate]rune, min_X, max_X, max_Y *int) {
	for i := *min_X-500; i< *max_X+500; i++{
		(*Cave)[Coordinate{i, *max_Y+2}]='#'
	}
	*min_X-=150
	*max_X+=150
	*max_Y+=1
}

func ForFunPrintResult(Cave *map[Coordinate]rune, min_X, max_X, max_Y *int){
	fmt.Println()
	for y := 0; y<=*max_Y+1; y++{
		for x := *min_X-1; x<=*max_X+1; x++{
			if (*Cave)[Coordinate{x,y}] == 0 {
				if y > *max_Y-2{
					if ((*Cave)[Coordinate{x+1,y}] == '#' && (*Cave)[Coordinate{x+1,y-1}] == '"'){
						(*Cave)[Coordinate{x,y}] = '"'
						fmt.Print("~")
					} else if ((*Cave)[Coordinate{x,y-1}] == '"' && (*Cave)[Coordinate{x+1,y-1}] == '#') {
						fmt.Print("~")
					} else{
						fmt.Print(".")
					}
				} else{
					fmt.Print(".")
				}
				
			} else if (*Cave)[Coordinate{x,y}] == '"'{
				fmt.Print("~")
			} else{
				fmt.Print(string((*Cave)[Coordinate{x,y}]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func RegolithReservoirPartOne(Cave map[Coordinate]rune, min_X, max_X, max_Y int){
	var sand int = 0

	out:
	for {
		newSand := Coordinate{500, 0}
		for {
			Cave[newSand]='#'-1
			if newSand.y +1 > max_Y{
				break out
			} else if Cave[Coordinate{newSand.x, newSand.y+1}] < '#'{
				newSand.y++
			} else if Cave[Coordinate{newSand.x-1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x--
			} else if Cave[Coordinate{newSand.x+1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x++
			} else{
				Cave[newSand] = 'o'
				sand++
				break
			}
			
			// Use to update picture every second (for small example)
			/*
			duration, _ := time.ParseDuration("100ms")
			time.Sleep(duration)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			ForFunPrintResult(&Cave, &min_X, &max_X, &max_Y)
			*/
			
		} 
	}
	// Just one Cave print at the end
	ForFunPrintResult(&Cave, &min_X, &max_X, &max_Y)

	fmt.Println("Part one result:", sand)
}

func RegolithReservoirPartTwo(Cave *map[Coordinate]rune, min_X, max_X, max_Y *int){
	var sand int = 0

	for {
		newSand := Coordinate{500, 0}
		if (*Cave)[newSand]=='o'{
			break
		}
		for {
			(*Cave)[newSand]='#'-1
			if (*Cave)[Coordinate{newSand.x, newSand.y+1}] < '#'{
				newSand.y++
			} else if (*Cave)[Coordinate{newSand.x-1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x--
			} else if (*Cave)[Coordinate{newSand.x+1, newSand.y+1}] < '#'{
				newSand.y++
				newSand.x++
			} else{
				sand++
				(*Cave)[newSand] = 'o'
				break
			}
			
			// Use to update picture every second (for small example)
			/*
			duration, _ := time.ParseDuration("100ms")
			time.Sleep(duration)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			ForFunPrintResult(Cave, min_X, max_X, max_Y)

			fmt.Println("Part two result:", sand)
			*/
		}
	}
	// Just one Cave print at the end
	ForFunPrintResult(Cave, min_X, max_X, max_Y)

	fmt.Println("Part two result:", sand)
}

func main(){

	Cave := make(map[Coordinate]rune)
	min_X, max_X, max_Y :=1000, 0, 0

	scanInput(&Cave, &min_X, &max_X, &max_Y)

	newCave := make(map[Coordinate]rune)
	for k,v := range Cave {
		newCave[k] = v
	}
	RegolithReservoirPartOne(newCave, min_X, max_X, max_Y)

	updateInput(&Cave, &min_X, &max_X, &max_Y)

	RegolithReservoirPartTwo(&Cave, &min_X, &max_X, &max_Y)

}