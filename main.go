package main

import (
	"BlockChainGo/BlockChain"
	"fmt"
	"log"
	"strconv"
)

func main() {
	chain := BlockChain.InitBlockChain()
	chain.AddBlock(BlockChain.BlockData{Name: "nasrallah"})
	chain.AddBlock(BlockChain.BlockData{Name: "adel"})
	chain.AddBlock(BlockChain.BlockData{Name: "nasrallah"})
	for _, block := range chain.Blocks {
		log.Printf("block.Hash >>> %x\n", block.Hash)
		log.Printf("block.PreviousHash >>> %x\n", block.PreviousHash)
		log.Printf("block.Data >>> %s", block.Data)
		pow := BlockChain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}

}
