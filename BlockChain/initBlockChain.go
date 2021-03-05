package BlockChain

func InitBlockChain() BlockChain {
	return BlockChain{[]*Block{Genesis()}}
}
