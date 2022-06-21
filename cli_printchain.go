package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := NewBlockChain()
	defer bc.db.Close()

	bci := bc.Interator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PreBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PreBlockHash) == 0 {
			break
		}
	}
}
