package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	wasm, err := ioutil.ReadFile("./basic.wasm")
	if err != nil {
		panic(err)
	}

	opts := &Options{}

	meterWasm, gasCost, _ := MeterWASM(wasm, opts)
	fmt.Println(meterWasm, gasCost)
}
