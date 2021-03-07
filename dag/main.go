// Copyright 2018 The godag Authors
// This file is part of the godag library.
//
// The godag library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The godag library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the godag library. If not, see <http://www.gnu.org/licenses/>.
//
// Join the discussion on https://godag.github.io
// See the Github source: https://github.com/garyyu/go-dag
//

package dag

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
)

func chainInitialize() map[string]*Block {

	//initial an empty chain
	chain := make(map[string]*Block)

	//add blocks

	AddBlock("Genesis", []string{})

	AddBlock("B", []string{"Genesis"})

	for i := 0; i < 8000; i++ {
		AddBlock(strconv.Itoa(rand.Int()), []string{"B"})
	}

	tips := FindTips(chain)
	tipsName := LTPQ(tips, true) // LTPQ is not relevant here, I just use it to get Tips name.
	AddBlock("Virtual", tipsName)

	return chain
}

func main() {

	var actual bytes.Buffer

	fmt.Println("\n- BlockDAG Algorithm Simulation - Algorithm 1: Selection of a blue set. -")

	chain := chainInitialize()

	fmt.Println("chainInitialize(): done. blocks=", len(chain)-1)

	CalcBlue(chain, 3, chain["Virtual"])

	// print the result of blue sets

	ltpq := LTPQ(chain, true)

	/*
		fmt.Print("blue set selection done. blue blocks = ")
		nBlueBlocks := 0
		actual.Reset()
		for _, name := range ltpq {
			block := chain[name]
			if IsBlueBlock(block)==true {
				if name=="Genesis" || name=="Virtual" {
					actual.WriteString(fmt.Sprintf("(%s).",name[:1]))
				}else {
					actual.WriteString(name+".")
				}

				nBlueBlocks++
			}
		}
	*/
	fmt.Println(actual.String(), "	total blue:", len(ltpq))

}
