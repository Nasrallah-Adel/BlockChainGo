package BlockChain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{
		Block:  b,
		Target: target,
	}
	return pow
}
func (pow *ProofOfWork) InitData(nonec int) []byte {
	dataByte := ConvertToArrayOfByte(pow.Block.Data)
	data := bytes.Join(
		[][]byte{
			pow.Block.PreviousHash,
			dataByte,
			ConvertToArrayOfByte(int64(nonec)),
			ConvertToArrayOfByte(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonec := 0
	for nonec < math.MaxInt64 {
		data := pow.InitData(nonec)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonec++
		}
	}
	fmt.Println()
	return nonec, hash[:]
}
