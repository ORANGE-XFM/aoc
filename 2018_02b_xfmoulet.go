package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_same(s1 string, s2 string) (bool, string) {
	// returns the common in s1 and s2 if differ by one or empty sring
	ndif:=0
	pos :=0
	s2runes := []rune(s2)
	for i,c := range s1 {
		if c!=s2runes[i] {
			ndif += 1
			pos = i
		}
	}

	return ndif==1, s1[:pos]+s1[pos+1:]
}

func process_file(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// read file line by line to an array
	ids := []string {}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		// not almost already seen ?
		for _, s2 := range ids {
			ok, sr := check_same(s, s2)
			if ok {
				fmt.Println("found",sr)
			}
		}
		ids = append(ids,s)
	}
}

func main() {
	process_file("2018_02_ex.data")
	process_file("2018_02.data")
}
