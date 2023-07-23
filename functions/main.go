package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Prefix(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func Postfix(postfix string) TransformFunc {
	return func(s string) string {
		return s + postfix
	}
}

func main() {
	var tes any
	tes = "Hello"
	fmt.Println(tes)
	tes = 1
	fmt.Println(tes)
	fmt.Println("Upper: ", transformString("Hello", UpperCase))
	fmt.Println("Prefix: ", transformString("Hello", Prefix("Prefix")))
	fmt.Println("Postfix: ", transformString("Hello", Postfix("Postfix")))
}
