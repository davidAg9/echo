package main

import (
	"fmt"

	"github.com/davidAg9/echo/chain"
)

func main() {
	echochain := chain.InitializeChain()

	echochain.AddBlock("First Block after genesis")
	echochain.AddBlock("Second Block after genesis")
	echochain.AddBlock("third Block after genesis")
	for _, block := range echochain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}
}
