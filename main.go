package main

import "fmt"

type MyFunc func(i int) string

func foo(i int) string {
	return fmt.Sprint(i)
}

type StringerFunc interface {
	Call(i int) string
}

func (f MyFunc) Call(i int) string {
	return fmt.Sprintf("%d called!", i)
}

func main() {
	var f MyFunc = foo
	fmt.Println(f(1))

	var sf StringerFunc = f
	fmt.Println(sf.Call(5))
}
