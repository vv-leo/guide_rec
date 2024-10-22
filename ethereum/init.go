package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var client *ethclient.Client

func init() {
	go func() {
		c, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}
		client = c
	}()
}
