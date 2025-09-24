package l1

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

func SolveL1_14(variable interface{}) {
	if t := reflect.TypeOf(variable); t != nil && t.Kind() == reflect.Chan {
		fmt.Printf("%v: %v\n", t, variable)
		return
	}

	switch v := variable.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: \"%s\"\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	default:
		fmt.Println("unknown type")
	}
}
