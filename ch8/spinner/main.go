// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 218.

// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"time"
)

var count int = 0
//JMWvar cache map[int]int

//!+
func main() {
    //JMWcache = make(map[int]int)
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d (count=%d)\n", n, fibN, count)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	count++
	if x < 2 {
		return x
	}
	//JMWif cache[x] != 0 {
		//JMWreturn cache[x]
	//JMW}
	//JMWfmt.Printf("fib(%d): doing return fib(%d) + fib(%d)\n", x, x-1, x-2)
	//JMWcache[x] = fib(x-1) + fib(x-2)
	//JMWreturn cache[x]
	return fib(x-1) + fib(x-2)
}

//!-
