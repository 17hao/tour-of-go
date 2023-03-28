package main

import (
	"fmt"
	"sort"
)

// p136 section 5.6
func main() {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"database":              {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures"},
		"programming languages": {"data structures"},
	}

	for i, item := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i, item)
	}
}

func topoSort(prereqs map[string][]string) []string {
	var order []string
	var visitAll func(items []string)
	seen := make(map[string]bool, 0)

	visitAll = func(items []string) {
		for _, item := range items {
			if _, ok := seen[item]; !ok {
				seen[item] = true
				visitAll(prereqs[item])
				order = append(order, item)
			}
		}
	}

	items := make([]string, 0)
	for k, _ := range prereqs {
		items = append(items, k)
	}
	sort.Strings(items)

	visitAll(items)

	return order
}
