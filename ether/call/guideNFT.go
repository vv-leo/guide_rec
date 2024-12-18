package call

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"web3/ether"
	"web3/ether/contracts"
)

const (
	privateKey1 = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey2 = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

// 调用上架nft测试方法
func ListNFT() {
	// 连接到以太坊客户端
	client, err := ether.EthConnector()
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
		return
	}

	// 设置合约地址
	contractAddress := viper.GetString("contracts.address.guideNFT")
	nft, err := contracts.NewGuideNFT(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		return
	}

	// 私钥应安全存储
	privateKeyHex := privateKey1
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("Error converting private key: %v", err)
		return
	}

	// 获取链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Failed to get network ID: %v", err)
		return
	}

	// 创建带有私钥的授权对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create auth object: %v", err)
		return
	}

	// 获取当前地址的余额
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("Failed to get balance: %v", err)
		return
	}
	log.Printf("Current balance: %s wei", balance.String())

	// 设定要上架的 token ID 和价格
	tokenId := big.NewInt(5)
	price := big.NewInt(77)

	// 获取预计的燃气费用
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to get gas price: %v", err)
		return
	}
	estimatedGasFee := new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice)
	log.Printf("Estimated gas fee: %s wei", estimatedGasFee.String())

	// 计算总费用
	totalCost := new(big.Int).Add(price, estimatedGasFee)
	if balance.Cmp(totalCost) < 0 {
		log.Println("Insufficient funds for transfer")
		return
	}

	// 调用合约的 ListNFT 方法
	_, err = nft.ListNFT(auth, tokenId, price)
	if err != nil {
		log.Printf("Failed to list NFT: %v", err)
		return
	}

	// 获取 token 信息
	infos, err := nft.TokenInfos(&bind.CallOpts{}, tokenId)
	if err != nil {
		log.Printf("Failed to get token infos: %v", err)
	} else {
		log.Printf("Token infos: %v", infos)
	}
}

func DeListNFT() {
	// 连接到以太坊客户端
	client, err := ether.EthConnector()
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
	}

	// 设置合约地址
	contractAddress := viper.GetString("contracts.address.guideNFT") // 确保这是正确的合约地址
	nft, err := contracts.NewGuideNFT(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
	}

	// 私钥应安全存储，以下为示例
	privateKeyHex := privateKey1
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("Error converting private key: %v", err)
	}

	// 获取链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Failed to get network ID: %v", err)
	}

	// 创建带有私钥的授权对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create auth object: %v", err)
	}
	// 调用合约的 ListNFT 方法
	tokenId := big.NewInt(5) // 设定要下架的 token ID
	res, err := nft.DeListNFT(auth, tokenId)
	if err != nil {
		log.Printf("Failed to delist NFT: %v", err)
	}
	log.Printf("Transaction hash: %v", res)

	infos, err := nft.TokenInfos(&bind.CallOpts{}, tokenId)

	if err != nil {
		log.Printf("Failed to get token infos: %v", err)
	} else {
		log.Printf("Token infos: %v", infos)
	}

}

func Buy() {
	// 连接到以太坊客户端
	client, err := ether.EthConnector()
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
	}

	// 设置合约地址
	contractAddress := viper.GetString("contracts.address.guideNFT") // 确保这是正确的合约地址
	nft, err := contracts.NewGuideNFT(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
	}

	// 私钥应安全存储，以下为示例
	privateKeyHex := privateKey2
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("Error converting private key: %v", err)
	}

	// 获取链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Failed to get network ID: %v", err)
	}

	// 创建带有私钥的授权对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create auth object: %v", err)
	}
	// 调用合约的 ListNFT 方法
	tokenId := big.NewInt(5)    // 设定要购买的 token ID
	auth.Value = big.NewInt(77) // 设定支付金额（需要大于或等于上架的价格）
	res, err := nft.BuyNFT(auth, tokenId)
	if err != nil {
		log.Printf("Failed to buy NFT: %v", err)
	}
	log.Printf("Transaction hash: %v", res)

	infos, err := nft.TokenInfos(&bind.CallOpts{}, tokenId)

	if err != nil {
		log.Printf("Failed to get token infos: %v", err)
	} else {
		log.Printf("Token infos: %v", infos)
	}

}
