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
	err := MakeMap(&mapSqrt, math.Sqrt)
	if err != nil {
		t.Error(err)
	}
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapSqrtTypeError(t *testing.T) {
	var mapSqrt func([]float64) []float64
	errIn := MakeMap(&mapSqrt, func(i int) float64 {
		return float64(i)
	})
	if errIn == nil {
		t.Error("errIn shoud be TypeError")
	}
	errOut := MakeMap(&mapSqrt, func(i float64) int {
		return int(i)
	})
	if errOut == nil {
		t.Error("errOut shoud be TypeError")
	}
}

func TestMakeMapTwice(t *testing.T) {
	var twice func([]string) []string
	err := MakeMap(&twice, func(i string) string { return i + i })
	if err != nil {
		t.Error(err)
	}
	r := twice([]string{"a", "bb", ""})
	rs := toS(r)
	es := `[]string{"aa", "bbbb", ""}`
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeFilter(t *testing.T) {
	var filterPlus func([]int) []int
	err := MakeFilter(&filterPlus, func(i int) bool { return i > 0 })
	if err != nil {
		t.Error(err)
	}
	r := filterPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "[]int{1, 10}"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeFilterTypeError(t *testing.T) {
	var filterPlus func([]int) []int
	errArg := MakeFilter(&filterPlus, func(i float64) bool { return i > 0 })
	if errArg == nil {
		t.Error("errArg shoud be TypeError")
	}
	errRet := MakeFilter(&filterPlus, func(i int) int { return 1 })
	if errRet == nil {
		t.Error("errRet shoud be TypeError")
	}
	errRetVoid := MakeFilter(&filterPlus, func(i int) {})
	if errRetVoid == nil {
		t.Error("errRet shoud be TypeError")
	}
}

func TestMakeSome(t *testing.T) {
	var hasOne func([]int) bool
	err := MakeSome(&hasOne, func(i int) bool { return i == 1 })
	if err != nil {
		t.Error(err)
	}
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
	err := MakeEvery(&everyPlus, func(i int) bool { return i > 0 })
	if err != nil {
		t.Error(err)
	}
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
	err := MakeReduce(&sum, func(r int, i int) int { return r + i })
	if err != nil {
		t.Error(err)
	}
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "15"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduce2(t *testing.T) {
	var sum func([]int) string
	err := MakeReduce(&sum, func(r string, i int) string { return r + fmt.Sprintf("%d", i) }, "0")
	if err != nil {
		t.Error(err)
	}
	r := sum([]int{1, 2, 3, 4, 5})
	es := "012345"
	if r != es {
		t.Error(r, "should be equal", es)
	}
}

func TestMakeReduceInit(t *testing.T) {
	var sum func([]int) int
	err := MakeReduce(&sum, func(r int, i int) int { return r + i }, 1)
	if err != nil {
		t.Error(err)
	}
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "16"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduceRight(t *testing.T) {
	var sum func([]int) int
	err := MakeReduceRight(&sum, func(r int, i int) int { return r - i })
	if err != nil {
		t.Error(err)
	}
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "-5"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeReduceRightInit(t *testing.T) {
	var sum func([]int) int
	err := MakeReduceRight(&sum, func(r int, i int) int { return r - i }, 20)
	if err != nil {
		t.Error(err)
	}
	r := sum([]int{1, 2, 3, 4, 5})
	rs := toS(r)
	es := "5"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapSqrtC(t *testing.T) {
	var mapSqrt func([]float64) []float64
	err := MakeMapC(&mapSqrt, math.Sqrt, 3)
	if err != nil {
		t.Error(err)
	}
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeFilterC(t *testing.T) {
	var filterPlus func([]int) []int
	err := MakeFilterC(&filterPlus, func(i int) bool { return i > 0 }, 3)
	if err != nil {
		t.Error(err)
	}
	r := filterPlus([]int{-10, -1, 0, 1, 10})
	rs := toS(r)
	es := "[]int{1, 10}"
	if toS(r) != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeSomeC(t *testing.T) {
	var hasOne func([]int) bool
	err := MakeSomeC(&hasOne, func(i int) bool { return i == 1 }, 3)
	if err != nil {
		t.Error(err)
	}
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
	err := MakeEveryC(&everyPlus, func(i int) bool { return i > 0 }, 3)
	if err != nil {
		t.Error(err)
	}
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
	err := MakeFirst(&mapSqrt, math.Sqrt)
	if err != nil {
		t.Error(err)
	}
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "1"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}

func TestMakeMapSqrtC_zero(t *testing.T) {
	var mapSqrt func([]float64) []float64
	err := MakeMapC(&mapSqrt, math.Sqrt, 0)
	if err != nil {
		t.Error(err)
	}
	r := mapSqrt([]float64{1, 4, 9})
	rs := toS(r)
	es := "[]float64{1, 2, 3}"
	if rs != es {
		t.Error(rs, "should be equal", es)
	}
}
