package enumerable_test

import (
	"enumerable"
	"fmt"
)

func ExampleMakeMap() {
	var twiceInt func([]int) []int
	enumerable.MakeMap(&twiceInt, func(i int) int { return i * 2 })
	fmt.Println(twiceInt([]int{1, 2, 3})) // [2 4 6]
}

func ExampleMakeFilter() {
	var filterOdd func([]int) []int
	enumerable.MakeFilter(&filterOdd, func(i int) bool { return i%2 == 1 })
	fmt.Println(filterOdd([]int{1, 2, 3})) // [1 3]
}

func ExampleMakeSome() {
	var hasOdd func([]int) bool
	enumerable.MakeSome(&hasOdd, func(i int) bool { return i%2 == 1 })
	fmt.Println(hasOdd([]int{1, 2, 3})) // true
	fmt.Println(hasOdd([]int{2, 4, 6})) // false
}

func ExampleMakeEvery() {
	var everyOdd func([]int) bool
	enumerable.MakeEvery(&everyOdd, func(i int) bool { return i%2 == 1 })
	fmt.Println(everyOdd([]int{1, 3, 5})) // true
	fmt.Println(everyOdd([]int{1, 2, 3})) // false
}

func ExampleMakeReduce() {
	var sumInt func([]int) int
	enumerable.MakeReduce(&sumInt, func(r int, i int) int { return r + i })
	fmt.Println(sumInt([]int{1, 2, 3})) // 6
}

func ExampleMakeReduceRight() {
	var minus func([]int) int
	enumerable.MakeReduceRight(&minus, func(r int, i int) int { return r - i })
	fmt.Println(minus([]int{1, 2, 3})) // 0
}
