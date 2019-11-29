package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_letters(msg string) (has2 bool, has3 bool) {
	seen_letters := map[rune]int{}
	for _, c := range msg {
		_, found := seen_letters[c]
		if found {
			seen_letters[c] += 1
		} else {
			seen_letters[c] = 1
		}
	}

	for _, v := range seen_letters {
		if v == 2 {
			has2 = true
		}
		if v == 3 {
			has3 = true
		}
	}
	// fmt.Println(msg, seen_letters, has2, has3)
	return
}

func process_file(filename string) {
	fmt.Println("hello")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// read file line by line
	n2 := 0
	n3 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hastwo, hasthree := check_letters(scanner.Text())
		if hastwo {
			n2 += 1
		}
		if hasthree {
			n3 += 1
		}
	}
	fmt.Println(n2, n3, n2*n3)
}

func main() {
	process_file("2018_02_ex.data")
	process_file("2018_02.data")
}
