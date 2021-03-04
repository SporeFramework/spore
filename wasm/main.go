package main

import (
	"fmt"

	"github.com/spore/wasm/metering"

	//"github.com/yyh1102/go-wasm-metering/toolkit"
	"io/ioutil"
)

func main() {
	wasm, err := ioutil.ReadFile("./basic.wasm")
	if err != nil {
		panic(err)
	}

	opts := &metering.Options{}

	meterWasm, gasCost, _ := metering.MeterWASM(wasm, opts)
	fmt.Println(meterWasm, gasCost)
}
