package BlockChain

////// CreateBlock  //////
func CreateBlock(data BlockData, previousHash []byte) Block {
	block := Block{
		Data:         data,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	pow := NewProof(&block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	//block.GenerateHash()
	return block
}
