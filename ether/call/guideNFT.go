package call

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"web3/ether"
	"web3/ether/contracts"
)

// 调用上架nft测试方法
func ListNFT() {
	// 连接到以太坊客户端
	client, err := ether.EthConnector()
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
	}

	// 设置合约地址
	contractAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3" // 确保这是正确的合约地址
	nft, err := contracts.NewGuideNFT(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
	}

	// 私钥应安全存储，以下为示例
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
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
	tokenId := big.NewInt(32) // 设定要上架的 token ID
	price := big.NewInt(1)    // 设定价格
	listNFT, err := nft.ListNFT(auth, tokenId, price)
	if err != nil {
		log.Printf("Failed to list NFT: %v", err)
	}

	// 输出成功信息
	fmt.Printf("NFT listed successfully with Token ID: %d\n", listNFT)
}
