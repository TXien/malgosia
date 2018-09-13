package main

import (
	"fmt"
	"lurcury/core/block"
	"lurcury/types"
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
