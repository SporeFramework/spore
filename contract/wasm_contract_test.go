package contract

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func Test_CreateWasmContract(t *testing.T) {
	eng, err := NewContractEngine()
	if err != nil {
		t.Error("Error constructing Wasm Contract Engine")
	}

	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		t.Errorf("Error opening wasm file: %s", err)
	}

	hash, gas, err := eng.CreateWasmContract(wasm)
	if err != nil {
		t.Errorf("Error creating Wasm contract: %s", err)
	}

	if gas != 1299 {
		t.Error("incorrect gas calculation")
	}

	if hex.EncodeToString(hash[:]) != "f5b012bbab7f165bc5eec4302f0952c1f3f8b601d4907cb8c5ae781a71821abb" {
		t.Error("Incorrect wasm contract digest")
	}

	t.Log(hex.EncodeToString(hash[:]), gas)
}

func Test_Call(t *testing.T) {

	eng, err := NewContractEngine()
	if err != nil {
		t.Error("Error constructing Wasm Contract Engine")
	}

	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		t.Errorf("Error opening wasm file: %s", err)
	}

	hash, gas, err := eng.CreateWasmContract(wasm)
	if err != nil {
		t.Errorf("Error creating Wasm contract: %s", err)
	}

	if gas != 1299 {
		t.Error("incorrect gas calculation")
	}

	result, gasCounter, err := eng.Call(hash, "increment")
	if gasCounter != 619 {
		t.Error("incorrect gas calculation")
	}

	resultValue := result.(int32)
	if resultValue != 1 {
		t.Error("Incorrect result from contract call returned")
	}

	t.Log(result, gasCounter)

}

func Test_CallMany(t *testing.T) {

	eng, err := NewContractEngine()
	if err != nil {
		t.Error("Error constructing Wasm Contract Engine")
	}

	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		t.Errorf("Error opening wasm file: %s", err)
	}

	hash, gas, err := eng.CreateWasmContract(wasm)
	if err != nil {
		t.Errorf("Error creating Wasm contract: %s", err)
	}

	if gas != 1299 {
		t.Error("incorrect gas calculation")
	}

	result, gasCounter, err := eng.Call(hash, "increment")
	if err != nil {
		t.Errorf("Error calling 'increment' function on Wasm contract: %s", err)
	}
	if gasCounter != 619 {
		t.Error("incorrect gas calculation")
	}

	resultValue := result.(int32)
	if resultValue != 1 {
		t.Error("Incorrect result from contract call returned")
	}

	var loops int32 = 1000

	var i int32
	for i = resultValue + 1; i < loops; i++ {
		res, gasUsed, err := eng.Call(hash, "increment")
		if err != nil {
			t.Errorf("Error calling 'increment' function on Wasm contract: %s", err)
		}
		if gasUsed != 619 {
			t.Errorf("Incorrect gas calculation at iteration %d", i)
		}
		if res.(int32) != i {
			t.Errorf("Incorrect result from contract call returned, was %d, expected %d", res.(int32), i)
		}
	}

	t.Log(result, gasCounter)

}
