package main

import (
	"fmt"
)

func gridval(serial int, x int, y int) int {
	rack_id := x + 10
	power_lvl := (rack_id*y + serial) * rack_id
	digit := (power_lvl / 100) % 10
	return digit - 5
}

func maxpower(matrix [301][301]int,size int) (int,int,int) {
		maxpwr, maxx, maxy := 0, 0, 0
	for i := 1; i <= 300-size+1; i++ {
		for j := 1; j <= 300-size+1; j++ {
			pwr := 0
			for di :=0 ; di<size ; di++ {
				for dj :=0 ; dj<size ; dj++ {
					pwr += matrix[i+di][j+dj]
				}
			}
			if pwr > maxpwr {
				maxx = i
				maxy = j
				maxpwr = pwr
			}
		}
	}
	return maxpwr,maxx,maxy
}

func main() {
	fmt.Println("gogogo")
	var matrix [301][301]int

	// Tests
	fmt.Println(gridval(8, 3, 5), 4)
	fmt.Println(gridval(57, 122, 79), -5)
	fmt.Println(gridval(39, 217, 196), 0)
	fmt.Println(gridval(71, 101, 153), 4)


	// Fill matrix
	//serial := 42
	serial := 9445

	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			matrix[i][j] = gridval(serial, i, j)
		}
	}

	for sz:=1;sz<=300;sz++ {
		maxpwr,maxx,maxy := maxpower(matrix,sz)
		fmt.Println("maxpwr",maxpwr,"x=",maxx,"y=", maxy, "sz=",sz)
	}
}
