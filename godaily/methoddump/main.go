package main

import (
	"fmt"
	"reflect"
)

type I1 interface {
	M()
	M1()
}
type I2 interface {
	M()
	M2()
}
type I interface {
	I1
	I2
}
type SI struct {
	I
}
type SI12 struct {
	I1
	I2
}

func main() {
	var si SI
	var si12 SI12
	DumpMethodSet(si)
	DumpMethodSet(si12)
}

func DumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}
	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}
	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
