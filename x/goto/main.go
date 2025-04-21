package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	now = time.Now()
	fmt.Println(impFib(50))
	fmt.Println(time.Since(now))

	now = time.Now()
	fmt.Println(fib(50))
	fmt.Println(time.Since(now))

	now = time.Now()
	fmt.Println(fib(50))
	fmt.Println(time.Since(now))

	fmt.Println("gib")
	now = time.Now()
	fmt.Println(gib(50))
	fmt.Println(time.Since(now))

	now = time.Now()
	fmt.Println(impFib(50))
	fmt.Println(time.Since(now))

	now = time.Now()
	fmt.Println(gotoFib(50))
	fmt.Println(time.Since(now))
}

func fib(n uint64) uint64 {
	return f(n, 0, 1)
}

func f(n, a, b uint64) uint64 {
	if n == 0 {
		return a
	}
	return f(n-1, b, (a + b))
}

func gib(n uint64) uint64 {
	a, b := uint64(0), uint64(1)
	return g(&n, &a, &b)
}

func g(n, a, b *uint64) uint64 {
	if *n == 0 {
		return *a
	}
	*n = *n - 1
	*a = *a + *b
	return g(n, b, a)
}

func impFib(n uint64) uint64 {
	a, b := uint64(0), uint64(1)
	for n--; n > 0; n-- {
		a, b = b, (a + b)
	}
	return b
}

func gotoFib(n uint64) uint64 {
	a, b := uint64(0), uint64(1)
	goto gFib
gFib:
	a, b = b, (a + b)
	n--
	if n == 0 {
		return a
	}
	goto gFib
}

func recFib(a uint64) uint64 {
	if a < 2 {
		return a
	}
	return recFib(a-1) + recFib(a-2)
}
