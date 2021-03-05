package BlockChain

type BlockData struct {
	Name string
}
type Block struct {
	Hash         []byte
	Data         BlockData
	PreviousHash []byte
}
type BlockChain struct {
	Blocks []*Block
}
