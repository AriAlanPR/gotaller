package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	p  = &v1          // has type *Vertex
)

func main() {
	//structs
	q := *p // *Vertex
	p.X = 2
	fmt.Println(v1, p, v2, v3, q)

	//arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// //slices
	var sl []int = primes[1:4]
	fmt.Println(sl)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	// //slice literals
	sli := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sli)

	// // For the array
	// // var a [10]int
	// // these slice expressions are equivalent:
	// // a[0:10]
	// // a[:10]
	// // a[0:]
	// // a[:]

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// // Extend its length.
	s = s[:4]
	printSlice(s)

	// // Drop its first two values.
	s = s[2:]
	printSlice(s)

	// // The zero value of a slice is nil.
	// // A nil slice has a length and capacity of 0 and has no underlying array.
	arr := make([]int, 5)     // len(a)=5, cap(a)=5
	arrb := make([]int, 0, 5) // len(b)=0, cap(b)=5

	printSlice(arr)
	printSlice(arrb)

	var arrc []int
	printSlice(arrc)

	// append works on nil slices.
	s = append(arrc, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	//range
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	// If you only want the index, you can omit the second variable.
	// for i := range pow
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// //interfaces
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "hello"
	// describe(i)

	// //functions as values
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X)*float64(v.X) + float64(v.Y)*float64(v.Y))
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// functions as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
