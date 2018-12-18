package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
	"log"
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
		//fmt.Println("s:",s,"a:",a,"b:",b1,"..",b2,horiz)

		if horiz {
			for x:=b1;x<=b2;x++ {
				grid[a][x] = '#'
			}
		} else {
			for y:=b1;y<=b2;y++ {
				grid[y][a] = '#'
			}
		}
	}

}


func (grid *Grid) flood(x int, y int) {

	//fmt.Println("flood",x,y)
	// source
	//grid[y][x] = '+'

	for grid[y][x] != '#' && grid[y][x] != '~' && y<H-1 {
		//grid[y][x]='|'
		y += 1
	}
	if y==H-1 {
		//fmt.Println("too low")
		return 
	}
	y -=1

	// extend left
	overflow := false
	x1 := x
	for x1>0 && grid[y][x1-1] == '.'  { 
		x1-- 
		if grid[y+1][x1-1] == '.' {
			grid.flood(x1-1,y)
			overflow = true
			goto next
		}
	}
	next:
	// extend right
	x2 := x
	for x2<W && grid[y][x2+1] == '.'  { 
		x2++ 
		if grid[y+1][x2+1] == '.' {
			grid.flood(x2+1,y)
			overflow = true
			goto next2
		}
	}
	next2:

	if overflow {
		return
	}

	// fill it ?
	// fmt.Println("fill it",y,x1,x2)
	if x1>0 && x2<W {
		for i:=x1;i<=x2;i++ {
			grid[y][i] = '~'
		}
	}
}
	
func (grid *Grid) print() {
	// output pgm file
	fmt.Println("P2")
	fmt.Println(W,H)
	fmt.Println(128)

	for y:=0;y<H;y++ { // 0..H
		for x:=0;x<W ; x++ { // 0..W
			fmt.Printf("%02d ",grid[y][x])
		}
		fmt.Print("\n")
	}
	log.Println("nb water",grid.count_water())
}

func (grid *Grid) count_water() int {
	nb := 0
	for y:=0;y<H;y++ {
		for x:=0;x<W;x++ {
			if grid[y][x]=='~' {
				nb += 1
			}
		}
	}
	return nb
}

func main() {
	var g Grid
	g.load("2018_17.data")
	prev := 0
	for i:=0;i<10000;i++ {
		g.flood(500,0)
		curr := g.count_water()
		if curr == prev {
			break
		}
		prev = curr
	}
	g.print()		
}
