package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/TwiN/go-color"
)

type BlockChain struct{
	blocks []*Block 
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash(){
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain*BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main(){
	chain := InitBlockChain()
	chain.AddBlock("FIRST BLOCK THROUGH GENESIS")
	chain.AddBlock("SECOND BLOCK THROUGH GENESIS")
	chain.AddBlock("THIRD BLOCK THROUGH GENESIS")

	for _, block := range chain.blocks{
		println(color.Red)
		fmt.Printf("Previous Hash -> %x\n", block.PrevHash)
		println(color.Reset)

		println(color.Purple)
		fmt.Printf("Data in Block -> %s\n", block.Data)
		println(color.Reset)

		println(color.Green)
		fmt.Printf("Hash -> %x\n", block.Hash)
		println(color.Reset)
	}
}

