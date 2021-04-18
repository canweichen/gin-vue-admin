package main

import "fmt"

var a A


type A struct {
	Endpoint string `yaml:"endpoint"`
}

func main(){
	C := A{}
	fmt.Println(C.Endpoint)
}
