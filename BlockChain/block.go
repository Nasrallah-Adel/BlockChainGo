package BlockChain

type BlockData struct {
	Name string
}
type Block struct {
	Hash         []byte
	Data         BlockData
	PreviousHash []byte
	Nonce        int
}
type BlockChain struct {
	Blocks []*Block
}
