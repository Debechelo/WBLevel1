package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func checkType1(value interface{}) {

	// type switch конструкция
	switch value.(type) {
	case int:
		fmt.Println("Value is int")
	case string:
		fmt.Println("Value is string")
	case bool:
		fmt.Println("Value is bool")
	case chan interface{}:
		fmt.Println("Value is chan interface {}")
	}
}

func checkType2(value interface{}) {
	typ := reflect.TypeOf(value)
	fmt.Printf("Value is %v\n", typ)
}

func main() {
	var a1 interface{} = true
	checkType1(a1)
	checkType2(a1)
	var a2 interface{} = 2
	checkType1(a2)
	checkType2(a2)
	var a3 interface{} = "true"
	checkType1(a3)
	checkType2(a3)
	var a4 interface{} = make(chan any)
	checkType1(a4)
	checkType2(a4)
}
