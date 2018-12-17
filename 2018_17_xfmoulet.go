package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

const H = 2000
const W = 1000

type Grid [H][W]uint8 

func (grid *Grid) load(filename string) {
	for i:=0;i<H;i++ {
		for j:=0;j<W;j++ {
			grid[i][j]='.'
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return 
	}
	defer file.Close() // -> not now !

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		horiz := s[0]=='y'
		ab := strings.Split(s,", ")
		a,_ := strconv.Atoi(ab[0][2:])

		b := strings.Split(ab[1],"..")
		b1,_ := strconv.Atoi(b[0][2:])
		b2,_ := strconv.Atoi(b[1])
		fmt.Println("s:",s,"a:",a,"b:",b1,"..",b2,horiz)

		if horiz {
			for x:=b1;x<=b2;x++ {
				grid[a][x] = '#'
				fmt.Println(x,a)
			}
		} else {
			for y:=b1;y<=b2;y++ {
				grid[y][a] = '#'
				fmt.Println(a,y)
			}
		}
	}

	// source
	grid[0][500] = '+'
}


func (grid *Grid) flood(x int, y int) {
	// flood fill down from position
	// fall from source to bottom or 
	grid[y][x] = '|'
	for grid[y][x] != '#' && y<H {
		grid[y][x]='|'
		fmt.Println("y",y)
		y += 1
	}
	y -=1

	// extend left
	for x1 := x; x1>0 && grid[y][x1] != '#' ; x1-- {
		grid[y][x1] = '~'
		fmt.Println("x1",x1)
	}
	// extend right
	for x2 := x; x2<W && grid[y][x2] != '#' ; x2++ {
		grid[y][x2] = '~'
		fmt.Println("x2",x2)
	}
}
	
func (grid *Grid) print() {
	for y:=0;y<25;y++ { // 0..H
		for x:=480;x<520 ; x++ { // 0..W
			fmt.Printf("%c ",grid[y][x])
		}
		fmt.Print("\n")
	}
}

func main() {
	var g Grid
	g.load("2018_17_ex.txt")
	g.print()
	g.flood(500,0)
	g.print()
}
