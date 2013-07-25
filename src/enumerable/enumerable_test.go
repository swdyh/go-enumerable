package enumerable

import (
	"fmt"
	"math"
	"testing"
)

func toS(a interface{}) string {
	return fmt.Sprintf("%#v", a)
}

func TestMakeMapSqrt(t *testing.T) {
	var mapSqrt func([]float64) []float64
	MakeMap(&mapSqrt, math.Sqrt)
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapTwice(t *testing.T) {
	var twice func([]string) []string
	MakeMap(&twice, func(i string) string { return i + i })
	r := twice([]string{"a", "bb", ""})
	rs := toS(r)
	es := `[]string{"aa", "bbbb", ""}`
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeFilter(t *testing.T) {
	var filterPlus func([]int) []int
	MakeFilter(&filterPlus, func(i int) bool { return i > 0 })
	r := filterPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "[]int{1, 10}"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeSome(t *testing.T) {
	var hasOne func([]int) bool
	MakeSome(&hasOne, func(i int) bool { return i == 1 })
	r := hasOne([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "true"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
	r2 := hasOne([]int{-10, -1, 0, 10})
	rs2 := toS(r2)
	es2 := "false"
	if rs2 != es2 {
		t.Error(rs2, "should be equal", es2)
	}
}

func TestMakeEvery(t *testing.T) {
	var everyPlus func([]int) bool
	MakeEvery(&everyPlus, func(i int) bool { return i > 0 })
	r := everyPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "false"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
	r2 := everyPlus([]int{10, 1, 3, 10})
	rs2 := toS(r2)
	es2 := "true"
	if rs2 != es2 {
		t.Error(rs2, "should be equal", es2)
	}
}

func TestMakeReduce(t *testing.T) {
	var sum func([]int) int
	MakeReduce(&sum, func(r int, i int) int { return r + i })
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "15"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduceInit(t *testing.T) {
	var sum func([]int) int
	MakeReduce(&sum, func(r int, i int) int { return r + i }, 1)
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "16"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduceRight(t *testing.T) {
	var sum func([]int) int
	MakeReduceRight(&sum, func(r int, i int) int { return r - i })
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "-5"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduceRightInit(t *testing.T) {
	var sum func([]int) int
	MakeReduceRight(&sum, func(r int, i int) int { return r - i }, 20)
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "5"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapSqrtC(t *testing.T) {
	var mapSqrt func([]float64) []float64
	MakeMapC(&mapSqrt, math.Sqrt, 3)
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeFilterC(t *testing.T) {
	var filterPlus func([]int) []int
	MakeFilterC(&filterPlus, func(i int) bool { return i > 0 }, 3)
	r := filterPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "[]int{1, 10}"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeSomeC(t *testing.T) {
	var hasOne func([]int) bool
	MakeSomeC(&hasOne, func(i int) bool { return i == 1 }, 3)
	r := hasOne([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "true"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
	r2 := hasOne([]int{-10, -1, 0, 10})
	rs2 := toS(r2)
	es2 := "false"
	if rs2 != es2 {
		t.Error(rs2, "should be equal", es2)
	}
}

func TestMakeEveryC(t *testing.T) {
	var everyPlus func([]int) bool
	MakeEveryC(&everyPlus, func(i int) bool { return i > 0 }, 3)
	r := everyPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "false"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
	r2 := everyPlus([]int{10, 1, 3, 10})
	rs2 := toS(r2)
	es2 := "true"
	if rs2 != es2 {
		t.Error(rs2, "should be equal", es2)
	}
}

func TestMakeFirst(t *testing.T) {
	var mapSqrt func([]float64) float64
	MakeFirst(&mapSqrt, math.Sqrt)
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "1"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapSqrtC_zero(t *testing.T) {
	var mapSqrt func([]float64) []float64
	MakeMapC(&mapSqrt, math.Sqrt, 0)
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}
