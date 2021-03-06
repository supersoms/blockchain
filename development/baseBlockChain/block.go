package main

import (
	"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
	"log"
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

//实现一个辅助函数将uint64转成[]byte类型
func uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
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
	/*var blockInfo []byte
	//1.拼装数据,将block.Data里面的数据打散然后追加到block.PrevHash后面
	blockInfo = append(blockInfo, uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)*/
	//TODO 以下代码对上面代码进行优化
	tmp := [][]byte{
		uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		uint64ToByte(block.TimeStamp),
		uint64ToByte(block.Difficulty),
		uint64ToByte(block.Nonce),
		block.Data,
	}
	//将二维切片数组链接起来,返回一个一维切片
	blockInfo := bytes.Join(tmp, []byte{})
	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
