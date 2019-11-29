package main

import (
	"fmt"
	"regexp"
	"bufio"
	"os"
	"strconv"
)

type Stars [][4]int

func readfile(filename string) Stars {
	var stars Stars 

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return Stars {}
	}
	defer file.Close()

	re := regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+),(-?\d+)$`)

	scanner := bufio.NewScanner(file)
	// read lines
	for scanner.Scan() {
		line := scanner.Text()
		res := re.FindStringSubmatch(line)[1:]
		if res != nil {
			var coord [4]int
			for i := range res {				
				val,err := strconv.Atoi(res[i])
				if err == nil { 
					coord[i] = val 
				}
			}
			stars = append(stars, coord)
		}
	}
	return stars
}
func manh(s1,s2 [4]int) {
	d := 0
	for dim:=0;dim<4;dim ++ 
		d += abs(s1[dim]-s2[dim])
	return d
}

func near(s1,s2 [4]int) {
	return manh(s1,s2)<=3
}

func (c Stars) has([4]int star) bool {
	for _,s := range c {
		if s==star { return true }
	}
	return false
}

// extend the constellation with star id in the stars
func find_constellation( id int, constellation, all_stars Stars ) []Stars {
	for i2,s := range stars {
		if near(stars[i],s) && !constellation.has{
			// append to constellation

		}
	}
	constellations = append(constellations, stars[0])
	// in turn, try to find a constellation to which to put current star 


	
}

func main() {
	stars := readfile("2018_25_ex1.txt")
	fmt.Println(stars)
}