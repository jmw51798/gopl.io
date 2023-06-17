// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
	"strings"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

// !+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	fmt.Println("============================================================")
	fmt.Println("topoSort: m=", m)
	fmt.Println("============================================================\n")
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	spacesStr := "                              "
	indentSpaces := 0

	visitAll = func(items []string) {
		indent := spacesStr[:indentSpaces]
		fmt.Printf("%svisitAll: items=[%s]\n", indent, strings.Join(items, ", "))
		for _, item := range items {
			if !seen[item] {
				indentSpaces += 2
				fmt.Printf("%sitem '%s' not seen, calling visitAll([%s])\n\n", indent, item, strings.Join(m[item], ", "))
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
				fmt.Printf("%safter visitAll() item=%s:\n", indent, item)
				fmt.Printf("%sorder=[%s]\n", indent, strings.Join(order, ", "))
				indentSpaces -= 2
			} else {
				fmt.Printf("%sitem '%s' already seen\n", indent, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Println("============================================================")
	fmt.Printf("keys=[%s]\n", strings.Join(keys, ", "))
	fmt.Println("============================================================\n")

	sort.Strings(keys)
	fmt.Println("============================================================")
	fmt.Printf("sorted keys=[%s]\n", strings.Join(keys, ", "))
	fmt.Println("============================================================\n")

	visitAll(keys)
	return order
}

//!-main
