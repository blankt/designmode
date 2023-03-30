package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBB")
	index := bytes.IndexByte(path, '/')

	a := path[:index]
	b := path[index+1:]
	fmt.Println(string(a))
	fmt.Println(string(b))

	//由于共享内存 所以a改变会改变b
	//a = append(a, "index"...)
	//fmt.Println(string(a))
	//fmt.Println(string(b))

	//不改变共享内存的方法 最后一个index指定a的cap
	a = path[:index:index]
	a = append(a, "index"...)
	fmt.Println(string(a))
	fmt.Println(string(b))
}
