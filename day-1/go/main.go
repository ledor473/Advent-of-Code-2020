package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	inventory := strings.Split(string(b), "\n\n")

	calories := make([]int, len(inventory))
	for i, inv := range inventory {
		items := strings.Split(inv, "\n")
		c := 0
		for _, item := range items {
			item, _ := strconv.Atoi(item)
			c += item
		}

		calories[i] = c
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Printf("Part 1: Most calories: %d\n", calories[0])

	fmt.Printf("Part 2: Sum of top 3: %d\n", sum(calories[0:3]))
}

func sum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
