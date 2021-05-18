package main

import (
	"fmt"
	parse2 "github.com/mangolau/the_ways_to_go/chapter_13/parse"
)

func main() {
	var examples = []string {
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n ", ex)
		nums, err := parse2.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
