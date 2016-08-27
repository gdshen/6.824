package main

import (
	"fmt"
)

func main() {
	//var kvs map[string]string

	kvs := make(map[string][]string)
	//append(kvs["hello"], "world")
	//kvs["hello"] = []string{"world"}

	//_ = kvs["hello"]
	//kvs["hello"] = "world"
	//if _, ok := kvs["hello"]; ok {
	//	fmt.Println("OK")
	//} else {
	//	fmt.Println("Not ok")
	//}
	//kvs["hello"] = make([]string, 0)
	//append(kvs["hello"], "world")
	//
	//fmt.Println(kvs)
	//kvs["hello"] = []string{"hello1, hello2"}
	kvs["hello"] = append(kvs["hello"], "hello3")

	fmt.Printf("%T\n%v\n", kvs, kvs)

}
