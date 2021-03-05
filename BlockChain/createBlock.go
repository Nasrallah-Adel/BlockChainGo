package BlockChain

////// CreateBlock  //////
func CreateBlock(data BlockData, previousHash []byte) Block {
	block := Block{
		Data:         data,
		PreviousHash: previousHash,
	}
	block.GenerateHash()
	return block
}
