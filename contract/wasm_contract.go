package contract

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	wasmtime "github.com/bytecodealliance/wasmtime-go"
	metering "github.com/sporeframework/spore/metering"

	"github.com/mathetake/gasm/hostfunc"
	"github.com/mathetake/gasm/wasi"
	"github.com/mathetake/gasm/wasm"
)

type ContractEngine struct {
	contracts  map[[32]byte]*wasmtime.Instance
	store      *wasmtime.Store
	gasCounter int64
}

func NewContractEngine() (*ContractEngine, error) {

	eng := &ContractEngine{
		contracts:  make(map[[32]byte]*wasmtime.Instance),
		store:      wasmtime.NewStore(wasmtime.NewEngine()),
		gasCounter: 0,
	}
	return eng, nil
}

func (engine *ContractEngine) gasConsumed(gas int64) {
	engine.gasCounter += gas
}

// CreateWasmContract creates a new contract
func (engine *ContractEngine) CreateWasmContract(wasm []byte) (sum [32]byte, gas uint64, err error) {

	sum = sha256.Sum256(wasm)
	if engine.contracts[sum] != nil {
		return sum, 0, errors.New("Contract already exists")
	}

	opts := &metering.Options{}
	meterWasm, gas, _ := metering.MeterWASM(wasm, opts)
	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(engine.store.Engine, meterWasm)
	check(err)

	item := wasmtime.WrapFunc(engine.store, engine.gasConsumed)
	// Instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := wasmtime.NewInstance(engine.store, module, []*wasmtime.Extern{item.AsExtern()})
	check(err)

	engine.contracts[sum] = instance
	return sum, gas, nil
}

// Call calls a wasm contract function
func (engine *ContractEngine) Call(contractID [32]byte, funcName string, args ...interface{}) (interface{}, int64, error) {

	// reset the gas counter
	defer func() {
		engine.gasCounter = 0
	}()

	instance := engine.contracts[contractID]

	if instance == nil {
		return nil, 0, errors.New("contract could not be found")
	}
	run := instance.GetExport(funcName).Func()
	result, err := run.Call(args...)

	// reset the gas counter
	return result, engine.gasCounter, err
}

// WasmTime test with wasmtime library
func WasmTime() {
	engine, err := NewContractEngine()
	if err != nil {
		panic(err)
	}

	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "in", "wasm", "basic.wasm"))
	//wasm, err := ioutil.ReadFile(path.Join("metering", "test", "expected-out", "wasm", "basic.wasm"))
	// wasm, err := ioutil.ReadFile("./add.wasm")
	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		panic(err)
	}

	id, gasUsed, _ := engine.CreateWasmContract(wasm)

	fmt.Println(gasUsed, id)
	// After we've instantiated we can lookup our `run` function and call
	// it.
	for i := 1; i <= 10; i++ {

		//result, err := Call(id, "addTwoNumbers", 42, 32)
		result, gas, err := engine.Call(id, "increment")

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
