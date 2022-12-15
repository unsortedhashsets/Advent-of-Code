package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"math"
	
)

type Coordinate struct{
	x,y int
}

type Sensor struct{
	coordinate Coordinate
	Beacon Coordinate
	distance int
}

func ManhattanDistance(x1, x2, y1, y2 *int) int{
	return int(math.Abs(float64(*x1-*x2))+math.Abs(float64(*y1-*y2)))
}

func scanInput(Sensors *[]Sensor, min_X, max_X, min_Y, max_Y *int) {
	readFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
	fs := bufio.NewScanner(readFile)

	var Sensor_X, Sensor_Y, Beacon_X, Beacon_Y int

	for fs.Scan() {
		_, _ = fmt.Sscanf(fs.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &Sensor_X, &Sensor_Y, &Beacon_X, &Beacon_Y)
		distance := ManhattanDistance(&Sensor_X, &Beacon_X, &Sensor_Y, &Beacon_Y)
		*Sensors = append(*Sensors, Sensor{Coordinate{Sensor_X, Sensor_Y}, Coordinate{Beacon_X, Beacon_Y}, distance})
		if Sensor_X-distance-2<*min_X{
			*min_X = Sensor_X-distance-2
		} 
		if Sensor_X+distance+2>*max_X{
			*max_X = Sensor_X+distance+2
		}
		if Sensor_Y-distance-2<*min_Y{
			*min_Y = Sensor_Y-distance-2
		}
		if Sensor_Y+distance+2>*max_Y{
			*max_Y = Sensor_Y+distance+2
		}
	}
}

// Iterate from x=0 to x={max_X}
// - Test each point through all sensors:
// 		AND:
// 		* If the point not a sensors's beacon (sensor.Beacon != Coordinate{x, y_input})
// 		* If the point is within the sensor's range (M_dis <= sensor.distance)
// 		* => count point as excluded and continue to the next point
// 		* Otherwise continue with the next sensor
func BeaconExclusionZonePartOne(Sensors *[]Sensor, max_X, min_X, max_Y, min_Y *int, y_input int){
	output := 0
	
	for x := *min_X; x <= *max_X; x++ {
		for _, sensor := range *Sensors{
			M_dis := ManhattanDistance(&x, &sensor.coordinate.x, &y_input, &sensor.coordinate.y)
			if (sensor.Beacon != Coordinate{x, y_input} && M_dis <= sensor.distance) {
				output++
				break
			}
		}

	}

	fmt.Println("Part one output", output)
}

// Iterate from y=0 to y={max_Y}
// 	 - Iterate from x=0 to x={max_X}
//		- Test each point through all sensors:
// 			* If the point distance to sensor is less than sensor's distance
// 			* => increase x on distance dif and break sensor's loop
// 			* Otherwise check next sensor
func BeaconExclusionZonePartTwo(Sensors *[]Sensor){
	min_X, min_Y := 0, 0
	max_X, max_Y := 4_000_000, 4_000_000

	for y := min_Y; y <= max_Y; y++ {
		out:
		for x := min_X; x <= max_X; x++ {
			testCoord := Coordinate{x, y}
			for _, sensor := range *Sensors{
				M_dis := ManhattanDistance(&testCoord.x, &sensor.coordinate.x, &testCoord.y, &sensor.coordinate.y)
				if M_dis <= sensor.distance {
					// Less efficient (2x)
					// x += (sensor.distance - M_dis)

					// Jump to sensor X coordinate
					x = sensor.coordinate.x
					// Add sensor distance with correction on Y distance
					x += sensor.distance - int(math.Abs(float64(sensor.coordinate.y-testCoord.y)))
					continue out
				}
			}
			fmt.Println("Part two output", x*max_X + y)
		}
	}	
}


func main(){
	Sensors := make([]Sensor, 0)

	min_X, max_X, min_Y, max_Y :=math.MaxInt, math.MinInt, math.MaxInt, math.MinInt

	scanInput(&Sensors, &min_X, &max_X, &min_Y, &max_Y)

	BeaconExclusionZonePartOne(&Sensors, &max_X, &min_X, &max_Y, &min_Y, 10)

	BeaconExclusionZonePartTwo(&Sensors)
}	