package main

import (
	"fmt"
	"strings"
)

// Array: a 1 indexed array
type Array []int

func (a Array) value(i int) int {
	return a[i-1]
}

func (a Array) swap(i int, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}

func HoarePartition(A Array, p int, r int) int {
	x := A.value(p)
	i := p - 1
	j := r + 1
	vars := func() string {
		return fmt.Sprintf(
			"A = %v, p = %v, r = %v, x = %v, i = %v, j = %v",
			A, p, r, x, i, j)
	}
	fmt.Println(vars())
	for {
		for {
			j = j - 1
			if A.value(j) <= x {
				break
			}
		}
		for {
			i = i + 1
			if A.value(i) >= x {
				break
			}
		}
		if i < j {
			A.swap(i, j)
		} else {
			fmt.Println(vars())
			return j
		}
		fmt.Println(vars())
	}
}

func QuickSort(A Array, p int, r int) {
	fmt.Printf("QuickSort(%v, %v, %v)\n", A, p, r)
	if p < r {
		q := HoarePartition(A, p, r)
		QuickSort(A, p, q)
		QuickSort(A, q+1, r)
	}
}

func main() {
	A := Array{13, 19, 9, 5, 12, 8, 7, 4, 11, 2, 6, 21}
	HoarePartition(A, 1, len(A))
	fmt.Println(strings.Repeat("-", 79))
	A = Array{13, 19, 9, 5, 12, 8, 7, 4, 11, 2, 6, 21}
	QuickSort(A, 1, len(A))
	fmt.Println(A)
}
