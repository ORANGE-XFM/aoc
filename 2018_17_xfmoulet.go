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

const WATER = '~'
const WET = '_'
const FREE = '\x00'
const BLOCK = '#'

type Grid [H][W]uint8 
var ymin, ymax int

func (grid *Grid) load(filename string) {
	for i:=0;i<H;i++ {
		for j:=0;j<W;j++ {
			grid[i][j]=FREE
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error reading file")
		return 
	}
	defer file.Close() // -> not now !

	scanner := bufio.NewScanner(file)
	ymin = 999999999
	ymax = -100
	for scanner.Scan() {
		s := scanner.Text()
		horiz := s[0]=='y'
		ab := strings.Split(s,", ")
		a,_ := strconv.Atoi(ab[0][2:])

		b := strings.Split(ab[1],"..")
		b1,_ := strconv.Atoi(b[0][2:])
		b2,_ := strconv.Atoi(b[1])
		//log.Println("s:",s,"a:",a,"b:",b1,"..",b2,horiz)

		if horiz {
			for x:=b1;x<=b2;x++ {
				grid[a][x] = BLOCK
			}
			if (a<ymin) { ymin = a }
			if (a>ymax) { ymax = a }
		} else {
			for y:=b1;y<=b2;y++ {
				grid[y][a] = BLOCK
			}
			if (b1<ymin) { ymin = b1 }
			if (b2>ymax) { ymax = b2 }
		}
	}
	log.Println("ymin",ymin,"ymax",ymax)

}

func (grid *Grid) canfall(x int, y int)  bool {
	c := grid[y][x]
	return c==FREE || c==WET
}
func (grid *Grid) extendLR(x int, y int)  bool {
	c := grid[y][x]
	return c != BLOCK
}

func (grid *Grid) flood(x int, y int) {

	log.Println("flood",x,y)
	// source
	//grid[y][x] = '+'

	for {
		for y<H-1 && grid[y][x]==FREE {
			grid[y][x]=WET
			y += 1
		}
		if y==H-1 || grid[y][x]==WET {
			//fmt.Println("too low")
			return 
		}
		y -=1

		overflow := false

		// extend left
		x1 := x
		for x1>0 && grid.extendLR(x1-1,y) { 
			x1-- 
			if grid.canfall(x1,y+1) {
				grid.flood(x1,y)
				overflow = true
				break
			}
		}
		// extend right
		x2 := x
		for x2<W && grid.extendLR(x2+1,y) { 
			x2++ 
			if grid.canfall(x2,y+1) {
				grid.flood(x2,y)
				overflow = true
				break
			}
		}

		if overflow {
			for i:=x1;i<=x2;i++ {
				grid[y][i] = WET
			}
			break
		} else {
			// fill it ?
			// fmt.Println("fill it",y,x1,x2)
			if x1>0 && x2<W {
				for i:=x1;i<=x2;i++ {
					grid[y][i] = WATER
				}
			}
		}

	}
}
	
func (grid *Grid) savepgm() {
	log.Println("outputting pgm ...")
	// output pgm file
	fmt.Println("P2")
	fmt.Println(W,ymax+5)
	fmt.Println(128)

	for y:=0;y<=ymax+5;y++ { // 0..H
		for x:=0;x<W ; x++ { // 0..W
			fmt.Printf("%02d ",grid[y][x])
		}
		fmt.Print("\n")
	}
}

func (grid *Grid) count_wet() int {
	nb := 0
	for y:=ymin;y<=ymax;y++ {
		for x:=0;x<W;x++ {
			if grid[y][x]==WATER || grid[y][x]==WET {
				nb += 1
			}
		}
	}
	return nb
}

func (grid *Grid) count_water() int {
	nb := 0
	for y:=ymin;y<=ymax;y++ {
		for x:=0;x<W;x++ {
			if grid[y][x]==WATER || (grid[y][x]==WET && grid[y-1][x]==WATER){
				nb += 1
			}
		}
	}
	return nb
}

func main() {
	var g Grid
	if false {
		g.load("2018_17_ex.txt")
	} else {
		g.load("2018_17.data")		
	}
	prev := 0
	for i:=0;i<1;i++ {
		g.flood(500,0)
		g[500][0] = '+'
		curr := g.count_water()
		if curr == prev {
			break
		}
		prev = curr
	}
	log.Println("nb wet",g.count_wet())
	log.Println("nb water",g.count_water())
	g.savepgm()

}
