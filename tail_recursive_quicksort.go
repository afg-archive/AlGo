package main

import (
	"fmt"
)

// Array: a 1 indexed array
type Array []int

func (a Array) value(i int) int {
	return a[i-1]
}

func (a Array) swap(i int, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}

func Partition(A Array, p int, r int) int {
	x := A.value(r)
	i := p - 1
	for j := p; j < r; j++ {
		if A.value(j) <= x {
			i++
			A.swap(i, j)
		}
	}
	A.swap(i+1, r)
	return i + 1
}

func TailRecursiveQuicksort(A Array, p int, r int) {
	fmt.Printf("TailRecursiveQuicksort(%v, %v, %v)\n", A, p, r)
	var q int
	for p < r {
		q = Partition(A, p, r)
		TailRecursiveQuicksort(A, p, q-1)
		p = q + 1
	}
}

func main() {
	A := Array{13, 19, 9, 5, 12, 8, 7, 4, 11, 2, 6, 21}
	TailRecursiveQuicksort(A, 1, len(A))
}
