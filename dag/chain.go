package dag

import (
	"fmt"
	"os"
)

//initial an empty chain
var chain map[string]*Block

func InitializeChain() {
	chain = make(map[string]*Block)
	AddBlock("Genesis", []string{})
	//tips := dag.FindTips(chain)
	//tipsName := dag.LTPQ(tips, true)
	//ChainAddBlock("Virtual", tipsName)
}

func InsertBlock(Name string) *Block {
	tips := FindTips(chain)
	keys := make([]string, 0, len(tips))
	for k := range tips {
		if chain[k] == nil {
			fmt.Println("PROBLEM")
		}
		keys = append(keys, k)
	}

	if len(keys) < 1 {
		fmt.Println("******** keys: ", keys)
	}

	block := AddBlock(Name, keys)
	fmt.Println("Chain size: ", len(chain))
	return block
}

var counter int = 0

func AddBlock(Name string, References []string) *Block {

	//create this block
	thisBlock := Block{Name, -1, -1, make(map[string]*Block), make(map[string]*Block), make(map[string]bool)}

	//add references
	for _, Reference := range References {
		prev, ok := chain[Reference]
		if ok {
			thisBlock.Prev[Reference] = prev
			prev.Next[Name] = &thisBlock
		} else {
			fmt.Println("chainAddBlock(): error! block reference invalid. block name =", Name, " references=", Reference)
			os.Exit(-1)
		}
	}
	//thisBlock.SizeOfPastSet = SizeOfPastSet(&thisBlock)
	counter++
	thisBlock.SizeOfPastSet = counter
	//add this block to the chain
	chain[Name] = &thisBlock
	return &thisBlock
}

func debugChain() {
	//debug
	tips := FindTips(chain)
	fmt.Println("tips: ", tips)
	tipsName := LTPQ(chain, true) // LTPQ is not relevant here, I just use it to get Tips name.
	AddBlock("Virtual", tipsName)

	CalcBlue(chain, 0, chain["Virtual"])
	ltpq := LTPQ(chain, true)
	fmt.Print("blue set selection done. blue blocks = ")
	fmt.Println(len(ltpq), ltpq)

	//fmt.Println(dag.Order(chain, 3))

	delete(chain, "Virtual")

}
