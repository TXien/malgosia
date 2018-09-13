
package main
 
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lurcury/core/transaction"
)
 
type T struct {
	A int64
	B float64
}

type TransactionJson struct{
        Nonce int64 //`json:"nonce"`
}

func TransactionBinaryEncode(trans transaction.TransactionJson)([]byte){
        buf := &bytes.Buffer{}
        err := binary.Write(buf, binary.BigEndian, trans)
        if err != nil {
                panic(err)
        }
	return buf.Bytes()
}

func TransactionBinaryDecode(target []byte)(trans transaction.TransactionJson){
	buf := bytes.NewBuffer(target)
	t := transaction.TransactionJson{}
	err := binary.Read(buf, binary.BigEndian, &t)
        if err != nil {
                panic(err)
        }
        return t
}

func main() {
	//t :=  T{A: 123}
	t := TransactionJson{Nonce:1}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
	//t = T{}
	//err = binary.Read(buf, binary.BigEndian, &t)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("%x %f", t.A, t.B)
}
