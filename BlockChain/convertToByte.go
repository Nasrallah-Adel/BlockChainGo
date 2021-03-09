package BlockChain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func ConvertToArrayOfByte(data interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
