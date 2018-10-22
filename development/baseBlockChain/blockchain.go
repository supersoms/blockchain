package main

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
