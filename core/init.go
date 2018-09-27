package main

import (
	"fmt"
	"lurcury/account"
	"lurcury/core/block"
	"lurcury/types"
	//"math/big"
)

func InitBlock()(types.BlockJson){
        b:=block.NewBlock("sue",
        0,
        "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
        0,
        0,
        "mogotisa",
        //"fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
        )
	fmt.Println(b)
	return b
}

func initAccount(core_arg types.CoreStruct, genesis types.BlockJson)(bool){
	fmt.Println("genesis:",len(genesis.Allocate))
	for i:=0; i<len(genesis.Allocate); i++{
		account.GenesisAccount(
			core_arg, 
			genesis.Allocate[i].Address, 
			genesis.Allocate[i].Amount,
		)
		fmt.Println(genesis.Allocate[i].Address,genesis.Allocate[i].Amount)
	}
	return true

}

func GenesisBlock()(types.BlockJson){
        genesis := types.BlockJson{
                Version: "sue",
                BlockNumber: 0,
                ParentHash: "",
                Nonce: 0,
                Timestamp: 0,
                ExtraData: "ka",
		Allocate: []types.AllocateStruct{
			{
				Address:"36b613fd8e2d15a61755836f4f70348968bd2478",
				Amount:"5000000000000000000000000000",
			},
		}, 
        }
	fmt.Println(genesis)
	return genesis

}
