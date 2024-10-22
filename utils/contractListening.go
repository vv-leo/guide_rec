package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 假设这是你的合约事件结构体
type YourContractEvent struct {
	address string
	gid     uint
	price   uint
}

// 启动一个goroutine来连接以太坊节点
func InitListen() {
	go func() {
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}
		//// 获取当前区块高度
		//header, err := client.HeaderByNumber(context.Background(), nil)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//startBlock := header.Number.Uint64()
		//
		//fmt.Printf("Starting to listen from block %d\n", startBlock)
		//
		//for {
		//	// 获取最新区块头
		//	header, err = client.HeaderByNumber(context.Background(), nil)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	currentBlock := header.Number.Uint64()
		//	// 检查是否有新区块
		//	if currentBlock > startBlock {
		//		for i := startBlock + 1; i <= currentBlock; i++ {
		//			block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//
		//			fmt.Printf("New block: %d\n", block.Number().Uint64())
		//			for _, tx := range block.Transactions() {
		//				fmt.Printf("  Transaction: %s\n", tx.Hash().Hex())
		//			}
		//		}
		//		startBlock = currentBlock
		//	}
		//
		//	// 等待一段时间再检查
		//	time.Sleep(2 * time.Second)
		//}
		// 在这里执行与以太坊节点的交互操作

		// 关闭连接
		//client.Close()
	}()
}
