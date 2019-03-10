// switch语句可以用于type-switch来判断某个interface 变量中实际存储的变量类型。
//判断数据类型(type),必须用onterface{}
package main

import (
	"fmt"
)

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x type:%T", i)
	case int:
		fmt.Printf("x type is int")
	case float64:
		fmt.Printf("x type is float64")
	case func(int) float64:
		fmt.Printf("x type is func(int)")
	case bool, string:
		fmt.Printf("x type is bool or string")
	default:
		fmt.Printf("x type is unknown")

	}
}