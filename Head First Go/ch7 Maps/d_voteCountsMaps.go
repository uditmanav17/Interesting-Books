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
	voteCounts := map[string]int{}
	for _, line := range lines {
		voteCounts[line]++
	}

	for candidate, votes := range voteCounts {
		fmt.Printf("%s: %d\n", candidate, votes)
	}

}
