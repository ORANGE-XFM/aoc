// very slow, should have used a circular linked list !
package main

import (
	"fmt"
)

func insert(s []int, i int, x int) []int {
	// insert after position i
	i %= len(s) + 1
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = x
	return s
}

func remove(a []int, i int) []int {
	// remove at position i
	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0 // or the zero value of T
	a = a[:len(a)-1]
	return a
}

func compute(NPLAYERS int, MARBLE int) {

	marbles := []int{0}
	pos := 0
	player := 0
	scores := make([]int, NPLAYERS)

	for marble := 1; marble <= MARBLE; marble++ {

		player += 1
		if player == NPLAYERS+1 {
			player = 1
		}

		if marble%23 == 0 {
			idx_rem := pos - 8
			if idx_rem < 0 {
				fmt.Println("plpl", marble, pos, idx_rem)
				idx_rem += len(marbles) // HERE
			}

			removed := marbles[idx_rem]
			//fmt.Println(marble,"removed",removed)
			marbles = remove(marbles, idx_rem)

			scores[player-1] += marble + removed
			pos = (idx_rem + 1) % len(marbles)
		} else {
			marbles = insert(marbles, pos+1, marble)
			pos = (pos + 2) % (len(marbles))
		}
		//fmt.Println(player,marbles,pos)
	}

	max := 0
	for _, n := range scores {
		if max < n {
			max = n
		}
	}
	//fmt.Println("scores :",scores)
	fmt.Println(NPLAYERS, "players", MARBLE, "marbles, max :", max, "\n")
}

func main() {

	// compute(9,25)
	// compute(10, 1618)
	compute(13, 7999)
	// compute(17,1104)
	// compute(21,6111)
	// compute(30,5807)

	compute(468, 71010)
	compute(468, 71010*100)

}
