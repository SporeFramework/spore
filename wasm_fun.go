package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"reflect"

	wasmtime "github.com/bytecodealliance/wasmtime-go"
	metering "github.com/sporeframework/spore/metering"

	"github.com/mathetake/gasm/hostfunc"
	"github.com/mathetake/gasm/wasi"
	"github.com/mathetake/gasm/wasm"
)

var contracts map[[32]byte]*wasmtime.Instance = make(map[[32]byte]*wasmtime.Instance)
var store = wasmtime.NewStore(wasmtime.NewEngine())

var gasTotal int64

func gasConsumed(gas int64) {
	fmt.Println("gas used: ", gas)
	gasTotal += gas
}

// CreateWasmContract creates a new contract
func CreateWasmContract(wasm []byte) (*wasmtime.Instance, [32]byte, uint64) {
	opts := &metering.Options{}
	meterWasm, gasCost, _ := metering.MeterWASM(wasm, opts)
	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(store.Engine, meterWasm)
	check(err)

	item := wasmtime.WrapFunc(store, gasConsumed)
	// Instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := wasmtime.NewInstance(store, module, []*wasmtime.Extern{item.AsExtern()})
	check(err)

	sum := sha256.Sum256(wasm)
	fmt.Printf("%x", sum)

	contracts[sum] = instance
	return instance, sum, gasCost
}

// Call calls a wasm contract function
func Call(contractID [32]byte, funcName string, args ...interface{}) (interface{}, int64, error) {

	// reset the gas counter
	gasTotal = 0
	instance := contracts[contractID]

	//meteringFunc := instance.GetExport("metering.usegas").Func()
	//fmt.Print(meteringFunc)
	run := instance.GetExport(funcName).Func()
	result, err := run.Call(args...)
	exp := instance.Exports()
	fmt.Println(exp)

	//check(err)
	// reset the gas counter
	return result, gasTotal, err
}

// WasmTime test with wasmtime library
func WasmTime() {
	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "in", "wasm", "basic.wasm"))
	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "expected-out", "wasm", "basic.wasm"))
	// wasm, err := ioutil.ReadFile("./add.wasm")
	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		panic(err)
	}

	_, id, gasUsed := CreateWasmContract(wasm)

	fmt.Println(gasUsed, id)
	// After we've instantiated we can lookup our `run` function and call
	// it.
	for i := 1; i <= 10; i++ {

		//result, err := Call(id, "addTwoNumbers", 42, 32)
		result, gas, err := Call(id, "increment")

		//run := instance.GetExport("increment").Func()
		//result, err := run.Call()
		fmt.Println("Total gas used: ", gas)
		check(err)
		fmt.Println(result)
	}

	// Serialize the module
	/*
		moduleBytes, err := module.Serialize()
		check(err)

		// deserialize and run
		module3, err := wasmtime.NewModuleDeserialize(store.Engine, moduleBytes)
		check(err)
		instance3, err := wasmtime.NewInstance(store, module3, []*wasmtime.Extern{item.AsExtern()})
		check(err)
		run := instance3.GetExport("increment").Func()
		result, err := run.Call()
		check(err)
		fmt.Println(result) // 42!
		fmt.Println("Total gas used: ", gasTotal)

		// instance 2
		run = instance2.GetExport("increment").Func()
		result, err = run.Call()
		check(err)
		fmt.Println(result) // 42!
	*/

}

// Gasm
func Gasm() {
	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "in", "wasm", "basic.wasm"))
	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "expected-out", "wasm", "basic.wasm"))
	buf, err := ioutil.ReadFile("./add.wasm")
	check(err)

	opts := &metering.Options{}

	meterWasm, gasCost, _ := metering.MeterWASM(buf, opts)
	fmt.Println(meterWasm, gasCost)

	mod, err := wasm.DecodeModule(bytes.NewBuffer(meterWasm))
	check(err)

	var gasTotal int64
	hostFunc := func(*wasm.VirtualMachine) reflect.Value {
		return reflect.ValueOf(func(gas int64) {
			fmt.Println("gas used: ", gas)
			gasTotal += gas
		})
	}

	builder := hostfunc.NewModuleBuilderWith(wasi.Modules)
	builder.MustSetFunction("metering", "usegas", hostFunc)
	vm, err := wasm.NewVM(mod, builder.Done())
	check(err)

	// fmt.Print(vm)

	ret, retTypes, err := vm.ExecExportedFunction("addTwoNumbers", 32, 43)
	check(err)

	fmt.Println(ret, retTypes)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
