# Go-Enumerable

[![Build Status](https://drone.io/github.com/swdyh/go-enumerable/status.png)](https://drone.io/github.com/swdyh/go-enumerable/latest)

Create enumerable functions(map, filter, some, every, reduce, reduceRight).

## Update

  * 2013-08-10 : Add type error.
  * 2013-07-25 : Add concurrent functions.
  * 2013-07-09 : First version

## Install

    go get github.com/swdyh/go-enumerable/src/enumerable

## Example

    package main
    
    import (
    	"fmt"
    	"github.com/swdyh/go-enumerable/src/enumerable"
    	"io/ioutil"
    	"net/http"
    )
    
    func main() {
    
    	// map for slice
    	var twiceInt func([]int) []int
    	enumerable.MakeMap(&twiceInt, func(i int) int { return i * 2 })
    	fmt.Println(twiceInt([]int{1, 2, 3})) // [2 4 6]

    	// type error
    	var twiceIntWrongType func([]int) []int
    	err := enumerable.MakeMap(&twiceIntWrongType, func(i int) string { return fmt.Sprintf("%d", i*2) })
    	fmt.Println(err) // TypeError: int != string
             
    	// filter for slice
    	var filterOdd func([]int) []int
    	enumerable.MakeFilter(&filterOdd, func(i int) bool { return i%2 == 1 })
    	fmt.Println(filterOdd([]int{1, 2, 3})) // [1 3]
    
    	// map for slice (concurrnt)
    	urls := []string{
    		"http://www.google.co.jp",
    		"http://www.yahoo.co.jp"}
    	var getAll func([]string) []string
    	enumerable.MakeMapC(&getAll, func(url string) string {
    		res, _ := http.Get(url)
    		b, _ := ioutil.ReadAll(res.Body)
    		return res.Status + " " + string(b)[0:50]
    	}, 2)
    	for _, v := range getAll(urls) {
    		fmt.Println(v)
    	}
    
    	// get first result (concurrnt)
    	var getFirst func([]string) string
    	enumerable.MakeFirst(&getFirst, func(url string) string {
    		res, _ := http.Get(url)
    		b, _ := ioutil.ReadAll(res.Body)
    		return url + " " + res.Status + " " + string(b)[0:50]
    	})
    	fmt.Println(getFirst(urls))
    }

https://github.com/swdyh/go-enumerable/blob/master/src/enumerable/example_makemap_test.go
https://github.com/swdyh/go-enumerable/blob/master/src/enumerable/enumerable_test.go


## Github

https://github.com/swdyh/go-enumerable

## License

The MIT License (MIT) Copyright (c) 2013 swdyh
