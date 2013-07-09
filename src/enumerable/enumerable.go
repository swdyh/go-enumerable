package enumerable

import (
	"reflect"
)

func MakeMap(fptr interface{}, f interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	out := fn.Type().Out(0)
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		s := reflect.MakeSlice(out, l, l)
		for i := 0; i < l; i++ {
			v := fv.Call([]reflect.Value{list.Index(i)})
			s.Index(i).Set(v[0])
		}
		return []reflect.Value{s}
	})
	fn.Set(v)
}

func MakeFilter(fptr interface{}, f interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	out := fn.Type().Out(0)
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		s := reflect.MakeSlice(out, 0, l)
		for i := 0; i < l; i++ {
			v := fv.Call([]reflect.Value{list.Index(i)})
			if v[0].Bool() {
				s = reflect.Append(s, list.Index(i))
			}
		}
		return []reflect.Value{s}
	})
	fn.Set(v)
}

func MakeSome(fptr interface{}, f interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		for i := 0; i < l; i++ {
			v := fv.Call([]reflect.Value{list.Index(i)})
			if v[0].Bool() {
				return []reflect.Value{reflect.ValueOf(true)}
			}
		}
		return []reflect.Value{reflect.ValueOf(false)}
	})
	fn.Set(v)
}

func MakeEvery(fptr interface{}, f interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		for i := 0; i < l; i++ {
			v := fv.Call([]reflect.Value{list.Index(i)})
			if !v[0].Bool() {
				return []reflect.Value{reflect.ValueOf(false)}
			}
		}
		return []reflect.Value{reflect.ValueOf(true)}
	})
	fn.Set(v)
}

func MakeReduce(fptr interface{}, f interface{}, iv ...interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		start := 0
		var r reflect.Value
		if len(iv) > 0 {
			r = reflect.ValueOf(iv[0])
		} else {
			r = list.Index(0)
			start = 1
		}
		for i := start; i < l; i++ {
			v := fv.Call([]reflect.Value{r, list.Index(i)})
			r = v[0]
		}
		return []reflect.Value{r}
	})
	fn.Set(v)
}

func MakeReduceRight(fptr interface{}, f interface{}, iv ...interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	v := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		list := in[0]
		l := list.Len()
		start := 0
		var r reflect.Value
		if len(iv) > 0 {
			r = reflect.ValueOf(iv[0])
		} else {
			r = list.Index(list.Len() - 1)
			start = 1
		}
		for i := (l - start - 1); i >= 0; i-- {
			v := fv.Call([]reflect.Value{r, list.Index(i)})
			r = v[0]
		}
		return []reflect.Value{r}
	})
	fn.Set(v)
}
