package enumerable

import (
	"fmt"
	"reflect"
)

type TypeError struct {
	Message string
}

func (te *TypeError) Error() string {
	return "TypeError: " + te.Message
}

func validateMapType(decType reflect.Type, impType reflect.Type) error {
	if err := assertArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := assertArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.Out(0).Elem(), impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateFilterType(decType reflect.Type, impType reflect.Type) error {
	if err := assertArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := assertArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0), decType.Out(0)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := assertBoolType(impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateSomeEveryType(decType reflect.Type, impType reflect.Type) error {
	if err := assertArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := assertArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := assertBoolType(decType.Out(0)); err != nil {
		return err
	}
	if err := assertBoolType(impType.Out(0)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	return nil
}

func validateReduceType(decType reflect.Type, impType reflect.Type, iv []interface{}) error {
	if err := assertArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := assertArgRetNum(impType, 2, 1); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0).Elem(), impType.In(1)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.Out(0), impType.In(0)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.Out(0), impType.Out(0)); err != nil {
		return err
	}
	if len(iv) > 0 {
		rType := reflect.TypeOf(iv[0])
		if err := assertTypeEqual(decType.Out(0), rType); err != nil {
			return err
		}
		if err := assertTypeEqual(impType.In(0), rType); err != nil {
			return err
		}
		if err := assertTypeEqual(impType.Out(0), rType); err != nil {
			return err
		}
	}
	return nil
}

func validateFirstType(decType reflect.Type, impType reflect.Type) error {
	if err := assertArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := assertArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := assertTypeEqual(decType.Out(0), impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func assertArgRetNum(funcType reflect.Type, numIn int, numOut int) error {
	if funcType.NumIn() != numIn {
		return &TypeError{fmt.Sprintf("arguments size error %v", funcType)}
	}
	if funcType.NumOut() != numOut {
		return &TypeError{fmt.Sprintf("return size error %v", funcType)}
	}
	return nil
}

func assertTypeEqual(decType reflect.Type, impType reflect.Type) error {
	if decType != impType {
		return &TypeError{fmt.Sprintf("%v != %v", decType, impType)}
	}
	return nil
}

func assertBoolType(t reflect.Type) error {
	if t.Kind() != reflect.Bool {
		return &TypeError{fmt.Sprintf("%v != bool", t)}
	}
	return nil
}
