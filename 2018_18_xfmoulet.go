package main
import (
	"fmt"
	"os"
	"bufio"
	"log"
)

const SZ = 50
const OPEN = '.'
const TREE = '|'
const LUMB = '#'

type Grid [SZ][SZ]uint8

func load (filename string) Grid {
	var g Grid

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error reading file")
		return g
	}
	defer file.Close() // -> not now !

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		s := scanner.Text()
		for j:=0;j<SZ;j++ {
			g[y][j]=s[j]
		}
		y += 1
	}

	return g
}

func (g Grid) print() {
	nb_tree,nb_lumb := 0,0
	for y:=0;y<SZ;y++ {
		for x:=0;x<SZ;x++ {
			switch g[y][x] {
				case TREE : nb_tree += 1
				case LUMB : nb_lumb += 1
			}
			fmt.Printf("%c ",g[y][x])
		}
		fmt.Println()
	}
	fmt.Println("Score :",nb_tree,"wood",nb_lumb,"lumber",nb_tree*nb_lumb,"total.")
}

func (g Grid) next_area(x int, y int) uint8 {
	xmin := x-1
	if xmin <0 { xmin = 0 }
	xmax := x+1
	if xmax ==SZ { xmax = SZ-1 }

	ymin := y-1
	if ymin <0 { ymin = 0 }
	ymax := y+1
	if ymax ==SZ { ymax = SZ-1 }

	nb_tree,nb_lumb := 0,0
	for i := xmin; i<=xmax ; i++ {
		for j := ymin; j<=ymax ; j++ {
			if i==x && j== y {continue} // skip center
			switch g[j][i] {
				case TREE : nb_tree += 1
				case LUMB : nb_lumb += 1
			}
		}
	}

	switch g[y][x]{
		case OPEN: if nb_tree>=3 {return TREE}
		case TREE: if nb_lumb>=3 {return LUMB}
		case LUMB: if nb_tree==0 || nb_lumb == 0 {return OPEN}
	}
	return g[y][x] // nothing happens
}

func (g Grid) next() Grid {
	var g2 Grid
	for i:=0;i<SZ;i++ {
		for j:=0;j<SZ;j++ {
			g2[j][i] = g.next_area(i,j)
		}
	}
	return g2
}

func main() {
	var g Grid
	g = load("2018_18.data")
	g.print()
	for i:=0;i<1000;i++ {
		g2 := g.next()
		if (g==g2) {break}
		g=g2
	}
	g.print()
}