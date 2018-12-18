package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func readfile(filename string) (records []string) {
	records = []string{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	sort.Strings(records)

	return
}

func calc_strat(records []string) {
	// newguard
	var lineguard = regexp.MustCompile(`^\[[0-9-]+ \d\d:\d\d\] Guard #(\d+) begins shift$`)
	var linere = regexp.MustCompile(`^\[[0-9-]+ \d\d:(\d\d)\] (.*)$`)

	// by guard : sum of sleeps
	guard_id := 0

	guards := make(map[int]int)
	guards_sleep := make(map[int]([60]int))

	start_sleep := 0
	for _, l := range records {
		mg := lineguard.FindStringSubmatch(l)
		if len(mg) > 0 {
			guard_id, _ = strconv.Atoi(mg[1])
		} else {
			m := linere.FindStringSubmatch(l)
			fmt.Println(m[1], m[2])
			if m[2] == "falls asleep" {
				start_sleep, _ = strconv.Atoi(m[1])
			} else {
				sleep, _ := strconv.Atoi(m[1])
				sleep -= start_sleep
				guard, _ := guards[guard_id]
				guards[guard_id] = guard + sleep

				v, _ := guards_sleep[guard_id]
				for i := 0; i < sleep; i++ {
					v[start_sleep+i] += 1
				}
				guards_sleep[guard_id] = v

				//fmt.Println(sleep, guard_id, guards, guards_sleep[guard_id])
			}
		}
	}

	// now get best guard
	best := -1
	bestindex := 0
	for i, guard := range guards {
		if guard > best {
			best = guard
			bestindex = i
		}
	}

	fmt.Println(best, bestindex, guards_sleep[bestindex])

	// now find most slept minute
	// now get best guard
	bestnb := -1
	bestmin := 0
	for min, nb := range guards_sleep[bestindex] {
		if nb > bestnb {
			bestmin = min
			bestnb = nb
		}
	}
	fmt.Println(bestmin * bestindex)

	// now find guard most asleep for this minute
	maxguard, maxmin, maxsleep := 0, 0, 0
	for guardid, sleep := range guards_sleep {
		for min := 0; min < 60; min++ {
			if sleep[min] > maxsleep {
				maxsleep = sleep[min]
				maxguard = guardid
				maxmin = min
			}
		}
	}
	fmt.Println(maxguard, maxmin*maxguard)
}

func main() {
	lines := readfile("2018_04.data")
	calc_strat(lines)

}
