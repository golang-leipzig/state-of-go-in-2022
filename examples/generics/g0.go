package main

import "fmt"

// Reverse returns a new slice, with elements in reverse order.
func Reverse[T any](vs []T) []T {
	var (
		n      = len(vs)
		result = make([]T, n)
	)
	for i, v := range vs {
		result[n-1-i] = v
	}
	return result
}

func main() {
	var (
		s  = []string{"a", "b", "c"}
		i  = []int{1, 2, 3}
		f  = []float64{1, 2, 3}
		rs = Reverse(s)
		ri = Reverse(i)
		rf = Reverse(f)
	)
	fmt.Printf("%v %T\n", rs, rs)
	fmt.Printf("%v %T\n", ri, ri)
	fmt.Printf("%v %T\n", rf, rf)
	// [c b a] []string
	// [3 2 1] []int
	// [3 2 1] []float64
}
