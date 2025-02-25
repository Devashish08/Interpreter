package evaluator

import (
	"fmt"
	"strings"

	"github.com/Devashish08/InterPreter-Compiler/object"
)

// Initialize built-in functions
func GetBuiltins() map[string]*object.Builtin {
	return map[string]*object.Builtin{
		"len":    {Fn: builtinLen},
		"first":  {Fn: builtinFirst},
		"last":   {Fn: builtinLast},
		"rest":   {Fn: builtinRest},
		"push":   {Fn: builtinPush},
		"puts":   {Fn: builtinPuts},
		"pop":    {Fn: builtinPop},
		"sum":    {Fn: builtinSum},
		"max":    {Fn: builtinMax},
		"min":    {Fn: builtinMin},
		"join":   {Fn: builtinJoin},
		"split":  {Fn: builtinSplit},
		"upper":  {Fn: builtinUpper},
		"lower":  {Fn: builtinLower},
	}
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return NewError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtinPuts(args ...object.Object) object.Object {
	for i, arg := range args {
		fmt.Print(arg.Inspect())
		if i < len(args)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	return NULL
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}
	return NULL
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}
	return NULL
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1)
		copy(newElements, arr.Elements[1:])
		return &object.Array{Elements: newElements}
	}
	return NULL
}

func builtinPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	newElements := make([]object.Object, len(arr.Elements)+1)
	copy(newElements, arr.Elements)
	newElements[len(arr.Elements)] = args[1]
	return &object.Array{Elements: newElements}
}

func builtinSplit(args ...object.Object) object.Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. got=%d, want=2", len(args))
	}

	str, ok := args[0].(*object.String)
	if !ok {
		return NewError("first argument to `split` must be STRING, got %s", args[0].Type())
	}

	delimiter, ok := args[1].(*object.String)
	if !ok {
		return NewError("second argument to `split` must be STRING, got %s", args[1].Type())
	}

	parts := strings.Split(str.Value, delimiter.Value)
	elements := make([]object.Object, len(parts))
	for i, part := range parts {
		elements[i] = &object.String{Value: part}
	}

	return &object.Array{Elements: elements}
}

func builtinJoin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. got=%d, want=2", len(args))
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return NewError("first argument to `join` must be ARRAY, got %s", args[0].Type())
	}

	delimiter, ok := args[1].(*object.String)
	if !ok {
		return NewError("second argument to `join` must be STRING, got %s", args[1].Type())
	}

	strs := make([]string, len(arr.Elements))
	for i, elem := range arr.Elements {
		if str, ok := elem.(*object.String); ok {
			strs[i] = str.Value
		} else {
			strs[i] = elem.Inspect()
		}
	}

	return &object.String{Value: strings.Join(strs, delimiter.Value)}
}

func builtinUpper(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}

	str, ok := args[0].(*object.String)
	if !ok {
		return NewError("argument to `upper` must be STRING, got %s", args[0].Type())
	}

	return &object.String{Value: strings.ToUpper(str.Value)}
}

func builtinLower(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}

	str, ok := args[0].(*object.String)
	if !ok {
		return NewError("argument to `lower` must be STRING, got %s", args[0].Type())
	}

	return &object.String{Value: strings.ToLower(str.Value)}
}

func builtinPop(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `pop` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length == 0 {
		return NewError("cannot pop from empty array")
	}

	return arr.Elements[length-1]
}

func builtinSum(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `sum` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return &object.Integer{Value: 0}
	}

	var sum int64
	for _, elem := range arr.Elements {
		if elem.Type() != object.INTEGER_OBJ {
			return NewError("array elements must be INTEGER, got %s", elem.Type())
		}
		sum += elem.(*object.Integer).Value
	}

	return &object.Integer{Value: sum}
}

func builtinMax(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `max` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return NULL
	}

	max := arr.Elements[0]
	if max.Type() != object.INTEGER_OBJ {
		return NewError("array elements must be INTEGER, got %s", max.Type())
	}
	maxVal := max.(*object.Integer).Value

	for _, elem := range arr.Elements[1:] {
		if elem.Type() != object.INTEGER_OBJ {
			return NewError("array elements must be INTEGER, got %s", elem.Type())
		}
		val := elem.(*object.Integer).Value
		if val > maxVal {
			maxVal = val
		}
	}

	return &object.Integer{Value: maxVal}
}

func builtinMin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return NewError("argument to `min` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return NULL
	}

	min := arr.Elements[0]
	if min.Type() != object.INTEGER_OBJ {
		return NewError("array elements must be INTEGER, got %s", min.Type())
	}
	minVal := min.(*object.Integer).Value

	for _, elem := range arr.Elements[1:] {
		if elem.Type() != object.INTEGER_OBJ {
			return NewError("array elements must be INTEGER, got %s", elem.Type())
		}
		val := elem.(*object.Integer).Value
		if val < minVal {
			minVal = val
		}
	}

	return &object.Integer{Value: minVal}
}
