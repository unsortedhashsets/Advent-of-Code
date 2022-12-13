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

func tailCatchUp(H Coordinate, T Coordinate) Coordinate {
	// Catch up diagonal R-U
	if (H.X - T.X > 1 && H.Y - T.Y > 0) || (H.X - T.X > 0 && H.Y - T.Y > 1) {
		T.X++
		T.Y++
	// Catch up diagonal R-D
	} else if (H.X - T.X > 1 && H.Y - T.Y < 0) || (H.X - T.X > 0 && H.Y - T.Y < -1) {
		T.X++
		T.Y--
	// Catch up diagonal L-U
	} else if (H.X - T.X < -1 && H.Y - T.Y > 0) || (H.X - T.X < 0 && H.Y - T.Y > 1) {
		T.X--
		T.Y++
	// Catch up diagonal L-D
	} else if (H.X - T.X < -1 && H.Y - T.Y < 0) || (H.X - T.X < 0 && H.Y - T.Y < -1) {
		T.X--
		T.Y--
	// Atomized moves
	} else {
		// Catch up X
		if H.X - T.X > 1{
			T.X++
		} else if H.X - T.X < -1 {
			T.X--
		}
		// Catch up Y
		if H.Y - T.Y > 1{
			T.Y++
		} else if H.Y - T.Y < -1 {
			T.Y--
		}
	}

	return T
}

func main(){

	tailTrack := make(map[Coordinate]bool)
	H := Coordinate{0, 0}
	T := Coordinate{0, 0}
	tailTrack[T] = true
	
	//Part two
	tailNineKnotTrack := make(map[Coordinate]bool)
	knots := make([]Coordinate, 8)
	for i := range knots {
		knots[i] = Coordinate{0, 0}
	}

	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

	fs := bufio.NewScanner(readFile)

	for fs.Scan() {
		direction := string(fs.Text())[0]
		steps_count, _ := strconv.Atoi(string(fs.Text())[2:])

		for s := steps_count; s > 0; s-- {

			switch direction {
			case 'R':
				H.X++
			case 'L':
				H.X--
			case 'U':
				H.Y++
			case 'D':
				H.Y--
			}	
			
			T = tailCatchUp(H, T)
			tailTrack[T] = true

			for i := range knots[:len(knots)] {
				if i == 0 {
					knots[i] = tailCatchUp(T, knots[i])
				} else {
					knots[i] = tailCatchUp(knots[i-1], knots[i])
				}
			}
			tailNineKnotTrack[knots[7]] = true
		}
	}
	
	readFile.Close()
	
	fmt.Println(len(tailTrack))
	fmt.Println(len(tailNineKnotTrack))
}

