package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func process_file(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)
	var linere = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	regs := [][5]int{}

	for scanner.Scan() {
		m := linere.FindStringSubmatch(scanner.Text())
		m2 := [5]int{0, 0, 0, 0, 0}
		for i := 0; i < 5; i++ {
			m2[i], _ = strconv.Atoi(m[i+1])
		}
		regs = append(regs, m2)
	}

	// count square inches and mark matrix
	matrix := [1000][1000]int{} // 0 if nothing, -1 if multiple
	nb := 0
	for _, l := range regs {
		id, x, y, w, h := l[0], l[1], l[2], l[3], l[4]
		for iy := y; iy < y+h; iy++ {
			for ix := x; ix < x+w; ix++ {
				if matrix[iy][ix] == 0 {
					matrix[iy][ix] = id
				} else {
					matrix[iy][ix] = -1
					nb += 1
				}
			}
		}
	}
	fmt.Println("total sq in", nb)

	for _, l := range regs {
		id, x, y, w, h := l[0], l[1], l[2], l[3], l[4]
		ok := true
		for iy := y; iy < y+h; iy++ {
			for ix := x; ix < x+w; ix++ {
				if matrix[iy][ix] == -1 {
					ok = false
				}
			}
		}
		if ok {
			fmt.Println("line OK", id)
		}
	}
}

func main() {
	process_file("2018_03_test.data")
	process_file("2018_03.data")
	//process_file("2018_02.data")
}
