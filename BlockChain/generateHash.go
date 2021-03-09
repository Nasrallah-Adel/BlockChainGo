package BlockChain

import (
	"bytes"
	"crypto/sha256"
)

////// GenerateHash //////
func (b *Block) GenerateHash() {
	dataByte := ConvertToArrayOfByte(b.Data)
	info := bytes.Join([][]byte{b.PreviousHash, dataByte}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
