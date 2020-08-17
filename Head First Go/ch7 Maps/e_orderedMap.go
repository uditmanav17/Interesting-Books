// maps are unordered if iterated over using for loop
// to access them in specific order everytime, we need different approach

package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"A": 54.452, "B": 98.8554}
	var names []string // create empty string slice
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.2f\n", name, grades[name])
	}
}
