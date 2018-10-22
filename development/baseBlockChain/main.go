package main

import (
	"fmt"
	"crypto/sha256"
)

//1.定义结构
type Block struct {
	//1.前区块哈希
	PrevHash []byte
	//2.当前区块哈希
	Hash []byte
	//3.数据
	Data []byte
}

//2.创建区块
func NewBlock(data string, prevBlock []byte) *Block {
	block := Block{
		PrevHash: prevBlock,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	block.SetHash()
	return &block
}

//3.生成哈希
func (block *Block) SetHash() {
	//1.拼装数据,将block.Data里面的数据打散然后追加到block.PrevHash后面
	blockInfo := append(block.PrevHash, block.Data...)
	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//4.引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("这是我的第一个BTC创世块！", []byte{})
}

//6.添加区块

func (bc *BlockChain) AddBlock(data string) {
	//如何获取前区块的哈希?
	lastBlock := bc.blocks[len(bc.blocks)-1]
	//获取最后一个区块的哈希值,把它作为前区块的哈希值
	prveHash := lastBlock.Hash
	//6.1创建新的区块
	block := NewBlock(data, prveHash)
	//6.2添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}

//7.重构代码

func main() {
	bc := NewBlockChain()
	bc.AddBlock("A向B转了50枚比特币!")
	bc.AddBlock("A又向B转了50枚比特币!")
	for i, block := range bc.blocks {
		fmt.Printf("====== 当前区块高度：%d ======\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n\n", block.Data)
	}
}
