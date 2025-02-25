package evaluator

import (
	"fmt"

	"github.com/Devashish08/InterPreter-Compiler/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func NewError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func IsError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR_OBJ
}

func IsTruthy(obj object.Object) bool {
	switch obj {
	case nil:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func NativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
