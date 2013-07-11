# Go-Enumerable

[![Build Status](https://drone.io/github.com/swdyh/go-enumerable/status.png)](https://drone.io/github.com/swdyh/go-enumerable/latest)

Create enumerable functions(map, filter, some, every, reduce, reduceRight).

## Install

    go get github.com/swdyh/go-enumerable/src/enumerable

## Example

    var twiceInt func([]int) []int
    enumerable.MakeMap(&twiceInt, func(i int) int { return i * 2 })
    fmt.Println(twiceInt([]int{1, 2, 3})) // [2 4 6]


    var filterOdd func([]int) []int
    enumerable.MakeFilter(&filterOdd, func(i int) bool { return i%2 == 1 })
    fmt.Println(filterOdd([]int{1, 2, 3})) // [1 3]

https://github.com/swdyh/go-enumerable/blob/master/src/enumerable/example_makemap_test.go
https://github.com/swdyh/go-enumerable/blob/master/src/enumerable/enumerable_test.go

## Github

https://github.com/swdyh/go-enumerable
