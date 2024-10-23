package call

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"web3/ether"
	"web3/ether/contracts"
)

const (
	privateKey      = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	accAddress      = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" //这里我使用transfer方法作为案例，所以需要一个接收转账地址
)

// SignTransaction 使用私钥对交易进行签名
//func SignTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
//	// 使用EIP-155标准计算交易的哈希
//	txHash := types.SignNewTx(tx.ChainID(), tx)
//
//	// 使用私钥对交易哈希进行签名
//	r, s, err := ecdsa.Sign(rand.Reader, privateKey, txHash[:])
//	if err != nil {
//		return nil, err
//	}
//
//	// 创建签名交易
//	signedTx := types.NewTx(&types.LegacyTx{
//		R:        r,
//		S:        s,
//		V:        big.NewInt(1).Add(big.NewInt(27).Quo(big.NewInt(0).SetUint64(privateKey.Curve.Params().N.BitLen()/8), big.NewInt(2)), big.NewInt(1)).Uint64(),
//		Nonce:    tx.Nonce(),
//		Gas:      tx.Gas(),
//		GasPrice: tx.GasPrice(),
//		To:       tx.To(),
//		Value:    tx.Value(),
//		Data:     tx.Data(),
//	})
//
//	return signedTx, nil
//}

// 示例的Signer函数
//var signerFn = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
//	privateKey, _ := crypto.HexToECDSA("your_private_key_hex_string")
//	return SignTransaction(tx, privateKey)
//}

func ListNFT() {
	client, err := ether.EthConnector()
	if err != nil {
		fmt.Println(err)
	}
	nft, err := contracts.NewGuideNFT(common.HexToAddress(contractAddress), client)

	//---------------------------------




	//---------------------------------

	transactOpts := bind.TransactOpts{
		From:     common.HexToAddress(accAddress),
		Signer:   SignerFn(mySignerFn)),
		GasLimit: 21000,
		Context:  nil,
	}
	listNFT, err := nft.ListNFT(&transactOpts, big.NewInt(1), big.NewInt(1))
	fmt.Println(listNFT)

	//privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	//privateKey, err := crypto.HexToECDSA(privateKeyHex)
	//if err != nil {
	//	log.Println("Error converting private key:", err)
	//	return
	//}
	//
	//// 发送者地址
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	//// 接收者地址
	//toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	//
	//// 交易金额
	//value := big.NewInt(1000000000000000000) // 1 Ether
	//
	//// 获取当前未使用的 nonce
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 设置 gas 价格和限制（可以根据实际情况调整）
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasLimit := uint64(21000)
	//
	//// 创建交易
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	//
	//// 对交易进行签名
	//signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 发送交易
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

}

type SignerFn func(types.Signer, common.Address, *types.Transaction) (*types.Transaction, error)

func mySignerFn(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	privateKeyHex := "your_private_key_hex_string"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
