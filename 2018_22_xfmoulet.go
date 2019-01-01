package main

import (
	"fmt"	
)

type Cave [][]int

// A region's erosion level is its geologic index 
// plus the cave system's depth, all modulo 20183
func erosion(level int, depth int) int {
	return (level+depth)%20183
}

// return erosion levels
func make_cave (depth int, target_x int, target_y int) *Cave {

	cave := make(Cave,target_y+1)
	for i:=0;i<=target_y;i++ {
		cave[i] = make([]int, target_x+1)		
	}

	cave[0][0] = erosion(0,depth)
	
	// If the region's Y coordinate is 0,
	// the geologic index is its X coordinate times 16807
	for x:=1;x<len(cave[0]);x++ {
		cave[0][x] = erosion(x*16807, depth)
	}
	// If the region's X coordinate is 0, 
	// the geologic index is its Y coordinate times 48271
	for y:=1;y<len(cave);y++ {
		cave[y][0] = erosion(y*48271, depth)
	}

	// Otherwise, the region's geologic index is the result of 
	// multiplying the erosion levels of the regions at X-1,Y and X,Y-1.
	for y:=1;y<=target_y;y++ {
		for x:=1;x<=target_x;x++ {
			cave[y][x] = erosion(cave[y-1][x]*cave[y][x-1], depth)
		}
	}
	
	cave[target_y][target_x] = erosion(0,depth)
	return &cave
}

func (cave *Cave ) print() {
	for _,line := range *cave {
		fmt.Println(line)
	}
}


func (cave *Cave) risk_level() int {
	risk := 0
	for _,line := range *cave {
		for _,v := range line {
			risk += v%3
		}
	}
	return risk
}

func main() {
	tst := make_cave(510,10,10)
	tst.print()
	fmt.Println(tst.risk_level())

	cave := make_cave(5355, 14, 796)
	//cave.print()
	fmt.Println(cave.risk_level())
}
