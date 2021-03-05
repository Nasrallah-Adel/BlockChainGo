package BlockChain

////// AddBlock //////
func (chain *BlockChain) AddBlock(data BlockData) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, &newBlock)
}
