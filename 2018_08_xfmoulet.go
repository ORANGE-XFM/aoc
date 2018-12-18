package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readfile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	//defer file.Close() -> not now !
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	return scanner
}

func printNums(scanner *bufio.Scanner) {
	// read file word by word
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		fmt.Println(x)
	}
}

func parseTree(scanner *bufio.Scanner) int {
	sum := 0
	scanner.Scan()
	nb_child, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	nb_meta, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < nb_child; i++ {
		sum += parseTree(scanner)
	}
	for i := 0; i < nb_meta; i++ {
		scanner.Scan()
		meta, _ := strconv.Atoi(scanner.Text())
		sum += meta
	}
	return sum
}

func parseTreeB(scanner *bufio.Scanner) int {
	scanner.Scan()
	nb_child, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	nb_meta, _ := strconv.Atoi(scanner.Text())

	child := make([]int, nb_child)
	for i := 0; i < nb_child; i++ {
		child[i] = parseTreeB(scanner)
	}

	if nb_child == 0 {
		sum := 0
		for i := 0; i < nb_meta; i++ {
			scanner.Scan()
			meta, _ := strconv.Atoi(scanner.Text())
			sum += meta
		}
		return sum
	} else {
		sum := 0
		for i := 0; i < nb_meta; i++ {
			scanner.Scan()
			index, _ := strconv.Atoi(scanner.Text())
			index -= 1
			if index >= 0 && index < nb_child {
				sum += child[index]
			}
		}
		return sum
	}
}

func main() {
	scanner := readfile("2018_08.data")
	fmt.Println(parseTree(scanner))
	scanner = readfile("2018_08.data")
	fmt.Println(parseTreeB(scanner))
}
