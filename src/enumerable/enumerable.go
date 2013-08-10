// Create enumerable functions(map, filter, some, every, reduce, reduceRight) for slice by reflection.
//
// References:
//   https://github.com/swdyh/go-enumerable
package enumerable

import (
	"reflect"
)

func MakeMap(fptr interface{}, f interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateMapType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	out := fn.Type().Out(0)
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
	return nil
}

func MakeFilter(fptr interface{}, f interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateFilterType(fn.Type(), fv.Type()); err != nil {
		return err
	}
	out := fn.Type().Out(0)
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
	return nil
}

func MakeSome(fptr interface{}, f interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateSomeEveryType(fn.Type(), fv.Type()); err != nil {
		return err
	}
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
	return nil
}

func MakeEvery(fptr interface{}, f interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateSomeEveryType(fn.Type(), fv.Type()); err != nil {
		return err
	}
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
	return nil
}

func MakeReduce(fptr interface{}, f interface{}, iv ...interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateReduceType(fn.Type(), fv.Type(), iv); err != nil {
		return err
	}
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
	return nil
}

func MakeReduceRight(fptr interface{}, f interface{}, iv ...interface{}) error {
	fn := reflect.ValueOf(fptr).Elem()
	fv := reflect.ValueOf(f)
	if err := validateReduceType(fn.Type(), fv.Type(), iv); err != nil {
		return err
	}
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
	return nil
}
