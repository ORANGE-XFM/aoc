package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func all_unique(arr []string) bool {
	uniques := map[string]bool{}
	for _, w := range arr {
		_, present := uniques[w]
		if present {
			return false
		} else {
			uniques[w] = true
		}
	}
	return true
}

func main() {
	fmt.Println("hello")
	file, err := os.Open("2017_04.data")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// read file line by line
	nlines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		is_unique := all_unique(words)
		fmt.Println(words, is_unique)
		if is_unique {
			nlines += 1
		}
	}
	fmt.Println(nlines)
}
