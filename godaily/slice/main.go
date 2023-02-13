package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	var sl1 []int
	var sl2 = []int{}

	fmt.Print("========基本区别=========\n")
	fmt.Printf("%v,len:%d,cap:%d,addr:%p\n", sl1, len(sl1), cap(sl1), &sl1)
	fmt.Printf("%v,len:%d,cap:%d,addr:%p\n", sl2, len(sl2), cap(sl2), &sl2)
	fmt.Printf("sl1==nil:%v\n", sl1 == nil)
	fmt.Printf("sl2==nil:%v\n", sl2 == nil)

	a1 := *(*[3]int)(unsafe.Pointer(&sl1))
	a2 := *(*[3]int)(unsafe.Pointer(&sl2))

	fmt.Print("========底层区别=========\n")
	fmt.Println(a1)
	fmt.Println(a2)

	type SliceDemo struct {
		Values []int
	}

	var s1 = SliceDemo{}
	var s2 = SliceDemo{[]int{}}
	bs1, _ := json.Marshal(s1)
	bs2, _ := json.Marshal(s2)

	fmt.Print("========序列化区别=========\n")
	fmt.Println(a1)
	fmt.Println(string(bs1))
	fmt.Println(string(bs2))
}

/*
========基本区别=========
[],len:0,cap:0,addr:0xc000004078
[],len:0,cap:0,addr:0xc000004090
sl1==nil:true
sl2==nil:false
========底层区别=========
[0 0 0]
[2876888 0 0]
========序列化区别=========
[0 0 0]
{"Values":null}
{"Values":[]}
*/
