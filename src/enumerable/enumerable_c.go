package enumerable

import (
	"reflect"
)

type listVal struct {
	Index int
	Val   reflect.Value
}

func runGoroutines(size int, fv reflect.Value) (chan listVal, chan listVal) {
	chin := make(chan listVal, size)
	chout := make(chan listVal, size)
	if size < 1 {
		size = 1
	}
	for i := 0; i < size; i++ {
		go func(i int, fv reflect.Value) {
			for in := range chin {
				chout <- listVal{in.Index, fv.Call([]reflect.Value{in.Val})[0]}
			}
		}(i, fv)
	}
	return chin, chout
}

func send(chin chan listVal, list reflect.Value) {
	size := list.Len()
	go func() {
		for i := 0; i < size; i++ {
			chin <- listVal{i, list.Index(i)}
		}
		close(chin)
	}()
}

func MakeMapC(fptr interface{}, f interface{}, gsize int) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateMapType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	out := fn.Type().Out(0)
	fr := func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		s := reflect.MakeSlice(out, l, l)
		chin, chout := runGoroutines(gsize, fv)
		send(chin, list)
		for i := 0; i < l; i++ {
			v := <-chout
			s.Index(v.Index).Set(v.Val)
		}
		return []reflect.Value{s}
	}
	fn.Set(reflect.MakeFunc(fn.Type(), fr))
	return nil
}

func MakeFilterC(fptr interface{}, f interface{}, gsize int) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateFilterType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	out := fn.Type().Out(0)
	fr := func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		s := reflect.MakeSlice(out, 0, l)
		chin, chout := runGoroutines(gsize, fv)
		send(chin, list)
		for i := 0; i < l; i++ {
			v := <-chout
			if v.Val.Bool() {
				s = reflect.Append(s, list.Index(v.Index))
			}
		}
		return []reflect.Value{s}
	}
	fn.Set(reflect.MakeFunc(fn.Type(), fr))
	return nil
}

func MakeSomeC(fptr interface{}, f interface{}, gsize int) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateSomeEveryType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	fr := func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		chin, chout := runGoroutines(gsize, fv)
		send(chin, list)
		for i := 0; i < l; i++ {
			v := <-chout
			if v.Val.Bool() {
				return []reflect.Value{reflect.ValueOf(true)}
			}
		}
		return []reflect.Value{reflect.ValueOf(false)}
	}
	fn.Set(reflect.MakeFunc(fn.Type(), fr))
	return nil
}

func MakeEveryC(fptr interface{}, f interface{}, gsize int) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateSomeEveryType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	fr := func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		chin, chout := runGoroutines(gsize, fv)
		send(chin, list)
		for i := 0; i < l; i++ {
			v := <-chout
			if !v.Val.Bool() {
				return []reflect.Value{reflect.ValueOf(false)}
			}
		}
		return []reflect.Value{reflect.ValueOf(true)}
	}
	fn.Set(reflect.MakeFunc(fn.Type(), fr))
	return nil
}

func MakeFirst(fptr interface{}, f interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateFirstType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	fr := func(in []reflect.Value) []reflect.Value {
		list := in[0]
		llen := list.Len()
		chout := make(chan listVal, llen)
		quit := make(chan bool, list.Len())
		for i := 0; i < llen; i++ {
			go func(i int, v reflect.Value) {
				select {
				case chout <- listVal{i, fv.Call([]reflect.Value{v})[0]}:
				case <-quit:
				}
			}(i, list.Index(i))
		}
		v := <-chout
		for i := 0; i < llen; i++ {
			quit <- true
		}
		return []reflect.Value{v.Val}
	}
	fn.Set(reflect.MakeFunc(fn.Type(), fr))
	return nil
}
