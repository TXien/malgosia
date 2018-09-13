package crypto

import (
	"encoding/hex"
	//eddsa "lurcury/crypto/eddsa"
	"testing"
	"reflect"
)

func CryptoTest(t *testing.T){
        check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }
	address := hex.EncodeToString(KeyToAddress([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c")))
	keccak := Keccak256([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c"))
	check("keyToAddress()", address, "3c01f961399c0c50a7ad37778eee8a20e1c1bb32")
	check("keccak256()", keccak,"7a5189acc8bb077a3e97489661da6facdbdbf805c8b69c2ece50135c2f0dbf74")
}

