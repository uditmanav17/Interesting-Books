// vote counts - difficult way
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	// slices to store names and counts
	var names []string
	var counts []int

	for _, line := range lines {
		matched := false
		for idx, name := range names { // loop each val in names
			if name == line {
				counts[idx]++
				matched = true
			}
		}
		if matched == false { // name didn't exist in names
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

}
