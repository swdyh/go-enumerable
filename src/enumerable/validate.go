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
	if err := validateArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := validateArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.Out(0).Elem(), impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateFilterType(decType reflect.Type, impType reflect.Type) error {
	if err := validateArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := validateArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0), decType.Out(0)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := validateBoolType(impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateSomeEveryType(decType reflect.Type, impType reflect.Type) error {
	if err := validateArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := validateArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := validateBoolType(decType.Out(0)); err != nil {
		return err
	}
	if err := validateBoolType(impType.Out(0)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	return nil
}

func validateReduceType(decType reflect.Type, impType reflect.Type) error {
	if err := validateArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := validateArgRetNum(impType, 2, 1); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0).Elem(), impType.In(1)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.Out(0), impType.In(0)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.Out(0), impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateFirstType(decType reflect.Type, impType reflect.Type) error {
	if err := validateArgRetNum(decType, 1, 1); err != nil {
		return err
	}
	if err := validateArgRetNum(impType, 1, 1); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.In(0).Elem(), impType.In(0)); err != nil {
		return err
	}
	if err := validateTypeEqual(decType.Out(0), impType.Out(0)); err != nil {
		return err
	}
	return nil
}

func validateArgRetNum(funcType reflect.Type, numIn int, numOut int) error {
	if funcType.NumIn() != numIn {
		return &TypeError{fmt.Sprintf("arguments size error %v", funcType)}
	}
	if funcType.NumOut() != numOut {
		return &TypeError{fmt.Sprintf("return size error %v", funcType)}
	}
	return nil
}

func validateTypeEqual(decType reflect.Type, impType reflect.Type) error {
	if decType != impType {
		return &TypeError{fmt.Sprintf("%v != %v", decType, impType)}
	}
	return nil
}

func validateBoolType(t reflect.Type) error {
	if t.Kind() != reflect.Bool {
		return &TypeError{fmt.Sprintf("%v != bool", t)}
	}
	return nil
}
