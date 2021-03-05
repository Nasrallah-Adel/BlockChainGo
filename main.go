package main

import (
	"BlockChainGo/BlockChain"
	"log"
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
	}
}
