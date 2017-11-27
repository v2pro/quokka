package runtime

import "reflect"

func Add(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) + typedRight
		case float64:
			return float64(typedLeft) + typedRight
		default:
			panic("can not add int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft + float64(typedRight)
		case float64:
			return typedLeft + typedRight
		default:
			panic("can not add float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("add does not support " + reflect.TypeOf(left).String())
	}
}

func Subtract(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) - typedRight
		case float64:
			return float64(typedLeft) - typedRight
		default:
			panic("can not subtract int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft - float64(typedRight)
		case float64:
			return typedLeft - typedRight
		default:
			panic("can not subtract float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("subtract does not support " + reflect.TypeOf(left).String())
	}
}
