package main

import (
	"crypto/sha256"
	"time"
)

//1.定义结构
type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.Merkel根(梅克尔根就是一个hash值,V4再介绍)
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数,也就是挖矿要找的数据
	Nonce uint64

	//a.当前区块哈希(正常比特币区块中是没有当前区块的哈希,为了实现方便而写)
	Hash []byte
	//b.数据
	Data []byte
}

//实现一个辅助函数将uint转成[]byte
func uint64ToByte(num uint64) []byte {
	return []byte{}
}

//2.创建区块
func NewBlock(data string, prevBlock []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlock,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, //随便填写无效值
		Nonce:      0, //同上
		Hash:       []byte{},
		Data:       []byte(data),
	}
	block.SetHash()
	return &block
}

//3.生成哈希
func (block *Block) SetHash() {
	var blockInfo []byte
	//1.拼装数据,将block.Data里面的数据打散然后追加到block.PrevHash后面
	blockInfo = append(blockInfo, uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)
	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
