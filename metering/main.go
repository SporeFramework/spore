package main

import (
	"fmt"

	wasmer "github.com/SporeFramework/spore/metering"

	//"github.com/yyh1102/go-wasm-metering/toolkit"
	"io/ioutil"
)

func main() {
	wasm, err := ioutil.ReadFile("./basic.wasm")
	if err != nil {
		panic(err)
	}

	opts := &wasmer.Options{}

	meterWasm, gasCost, _ := wasmer.MeterWASM(wasm, opts)
	fmt.Println(meterWasm, gasCost)
}
