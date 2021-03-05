package BlockChain

func Genesis() *Block {
	initBlock := CreateBlock(BlockData{Name: "Genesis"}, []byte{})
	return &initBlock
}
