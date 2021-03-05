package BlockChain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

////// GenerateHash //////
func (b *Block) GenerateHash() {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b.Data)
	if err != nil {
		log.Fatal(err)
	}
	info := bytes.Join([][]byte{b.PreviousHash, buf.Bytes()}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
