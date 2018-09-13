package params

import (
	//"fmt"
	"math/big"
) 

type ChainConfigStructure struct {
	Id int64
	Hash string
	V string
	Version *Version
}

type VersionData struct {
        Fee *big.Int
        FeeAddress string
        FeeToken string
        BlockSpeed int
        BlockTransaction int
        Consensus string
}

type Version struct {
        Sue *VersionData
}

func Chain()(*ChainConfigStructure){
	d := &VersionData{
		Fee : big.NewInt(10),
		FeeAddress: "gx",
		FeeToken: "def",
		BlockSpeed: 1,
	}
	v := &Version{
		Sue : d,
	}

	s := &ChainConfigStructure{
		Hash: "0x",
		Id: 101,
		V: "0",
		Version: v,
	}
	return s
}


/*
func main(){
	fmt.Println(Chain().hash)
}
*/