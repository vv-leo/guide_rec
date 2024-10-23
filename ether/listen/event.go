package listen

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

type NFTListedEvent struct {
	Recipient common.Address
	TokenId   *big.Int
	Price     *big.Int
}

func EventListen() {
	go func() {
		client, err := ethclient.Dial("ws://127.0.0.1:8545")
		if err != nil {
			log.Printf("Failed to connect to Ethereum client: %v", err)
		}

		contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

		// 读取 .abi 文件
		data, err := ioutil.ReadFile("./ether/abi/GuideNFT.abi")
		if err != nil {
			log.Printf("Failed to read ABI file: %v", err)
		}

		parsedABI, err := abi.JSON(strings.NewReader(string(data)))
		if err != nil {
			log.Printf("Failed to parse contract ABI: %v", err)
		}

		// 确保 NFTListed 事件存在
		nftListedEvent := parsedABI.Events["NFTListed"]
		fmt.Printf("NFTListed Event ID: %s\n", nftListedEvent.ID.Hex())

		// 监听事件
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			Topics:    [][]common.Hash{{nftListedEvent.ID}}, // 设置主题
		}

		// 使用 SubscribeFilterLogs 进行实时监听
		logsCh := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
		if err != nil {
			log.Printf("Failed to subscribe to logs: %v", err)
		}

		// 处理日志
		for {
			select {
			case err := <-sub.Err():
				log.Printf("Error from subscription: %v", err)
				return
			case vLog := <-logsCh:
				fmt.Printf("Log received: %+v\n", vLog) // 打印日志信息

				event := NFTListedEvent{}

				err := parsedABI.UnpackIntoInterface(&event, "NFTListed", vLog.Data)
				if err != nil {
					log.Printf("Failed to unpack log: %v", err)
					continue
				}

				fmt.Printf("Event received: To: %s, TokenId: %s, Price: %s\n", event.Recipient.Hex(), event.TokenId.String(), event.Price.String())
			}
		}
	}()
}
