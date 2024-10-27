package ether

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// EthConnector 封装了与以太坊节点的连接逻辑
func EthConnector() (*ethclient.Client, error) {
	c, err := ethclient.Dial(viper.GetString("ethereum.addr.https"))
	if err != nil {
		fmt.Println("Failed to connect to Ethereum client:", err)
		return nil, err // 返回错误，而不是使用 log.Fatal
	}
	return c, nil // 返回建立的连接
}

var ethCon *ethclient.Client // 假设这是全局变量，用于存储连接
