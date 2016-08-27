package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyCustomType struct {
	Hello string
	World string
}

func main() {
	fileName := "json.txt"
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Open file failed")
	}

	enc := json.NewEncoder(f)

	enc.Encode(MyCustomType{"hi", "sgd"})
	enc.Encode(MyCustomType{"hi2", "sgd2"})

	f.Close()

	// decode usage

	f, err = os.Open("decode")
	if err != nil {
		fmt.Println("Decode phrase open file failed")
	}
	dec := json.NewDecoder(f)

	//var myCustomData MyCustomType

	//for {
	//	if err = dec.Decode(&myCustomData); err == io.EOF {
	//		break
	//	} else if err != nil {
	//		fmt.Println("Decode phrase decode failed")
	//	}
	//	fmt.Println("hello -> ", myCustomData.Hello, "; world -> ", myCustomData.World)
	//}
	//f.Close()

	var myCustom []string

	dec.Decode(&myCustom)
	// note the difference between json objects and array of json objects
	fmt.Printf("%v\n", myCustom)
	f.Close()

}
