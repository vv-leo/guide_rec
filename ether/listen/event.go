package listen

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"web3/model"
)

type NFTListedEvent struct {
	Recipient common.Address
	TokenId   *big.Int
	Price     *big.Int
}

type EventNFTListed struct {
	Recipient common.Address
	TokenId   *big.Int
	Price     *big.Int
}

type EventNFTDelisted struct {
	Operator common.Address
	TokenId  *big.Int
}

type EventNFTBought struct {
	Buyer     common.Address
	Seller    common.Address
	TokenId   *big.Int
	Price     *big.Int
	TimeStamp *big.Int
}

var guideDao = model.NewGuideDao()
var buyRecordDao = model.NewBuyRecord()

func EventListen() {
	go func() {

		addr := viper.GetString("ethereum.addr.ws")
		client, err := ethclient.Dial(addr)
		if err != nil {
			log.Printf("Failed to connect to Ethereum client: %v", err)
		}

		contractAddress := common.HexToAddress(viper.GetString("contracts.address.guideNFT"))

		// 读取 .abi 文件
		data, err := ioutil.ReadFile("./ether/abi/GuideNFT.abi")
		if err != nil {
			log.Printf("Failed to read ABI file: %v", err)
		}

		parsedABI, err := abi.JSON(strings.NewReader(string(data)))
		if err != nil {
			log.Printf("Failed to parse contract ABI: %v", err)
		}

		//// 确保 NFTListed 事件存在
		nftListedEvent := parsedABI.Events["NFTListed"]
		nftDeListedEvent := parsedABI.Events["NFTDelisted"]
		nftBoughtEvent := parsedABI.Events["NFTBought"]

		// 监听事件
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			Topics:    [][]common.Hash{{nftListedEvent.ID, nftDeListedEvent.ID, nftBoughtEvent.ID}}, // 设置主题
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

				switch {
				case vLog.Topics[0].Hex() == parsedABI.Events["NFTListed"].ID.Hex():
					event := struct {
						Recipient common.Address
						TokenId   *big.Int
						Price     *big.Int
					}{}

					// 从 Topics 和 Data 中解析数据
					tokenId := new(big.Int).SetBytes(vLog.Topics[2][12:]) // 解析 tokenId
					price := new(big.Int).SetBytes(vLog.Data[0:32])       // 解析价格

					event.Recipient = common.HexToAddress(vLog.Topics[1].Hex())
					event.TokenId = tokenId
					event.Price = price

					log.Printf("NFTListed Event: %+v", event)
					//更新库上架状态及价格
					guideDao.List(event.TokenId.String(), event.Price.String())

				case vLog.Topics[0].Hex() == parsedABI.Events["NFTDelisted"].ID.Hex():
					event := EventNFTDelisted{}

					// 从 Topics 和 Data 中解析数据
					tokenId := new(big.Int).SetBytes(vLog.Topics[2][12:]) // 解析 tokenId

					event.Operator = common.HexToAddress(vLog.Topics[1].Hex())
					event.TokenId = tokenId

					log.Printf("NFTDelisted Event: %+v", event)

					//更新库下架状态
					guideDao.DeList(event.TokenId.String())

				case vLog.Topics[0].Hex() == parsedABI.Events["NFTBought"].ID.Hex():
					event := EventNFTBought{}
					// 从 Topics 和 Data 中解析数据
					tokenId := new(big.Int).SetBytes(vLog.Topics[3][12:]) // 解析 tokenId
					price := new(big.Int).SetBytes(vLog.Data[0:32])       // 解析价格
					time := new(big.Int).SetBytes(vLog.Data[32 : 32+32])  //解析时间戳

					event.Buyer = common.HexToAddress(vLog.Topics[1].Hex())
					event.Seller = common.HexToAddress(vLog.Topics[2].Hex())
					event.TokenId = tokenId
					event.Price = price
					event.TimeStamp = time

					log.Printf("NFTBought Event: %+v", event)

					//交易记录写入数据库
					record := model.BuyRecord{
						GuideId: event.TokenId.String(),
						Buyer:   event.Buyer.String(),
						Seller:  event.Seller.String(),
						Price:   event.Price.String(),
						Time:    event.TimeStamp.Int64(),
					}

					buyRecordDao.CreateBuyRecord(record)
					guideDao.UpdateOwner(event.TokenId.String(), event.Buyer.String())

				default:
					log.Printf("Unknown event received: %+v", vLog)
				}
			}
		}

	}()
}
