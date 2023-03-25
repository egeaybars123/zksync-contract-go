package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/zksync-sdk/zksync2-go"
)

var rpcClientAddress = "https://goerli.infura.io/v3/" + GetEnv("API_KEY")

// use this function only for string outputs
func GetEnv(env string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv(env)
	return val
}

func InitWallet() (*zksync2.Wallet, error) {
	ethereumSigner, err := zksync2.NewEthSignerFromMnemonic(GetEnv("MNEMONIC_KEY"), 280)
	if err != nil {
		return nil, err
	}

	zkSyncProvider, err := zksync2.NewDefaultProvider("https://zksync2-testnet.zksync.dev")
	if err != nil {
		return nil, err
	}

	wallet, err := zksync2.NewWallet(ethereumSigner, zkSyncProvider)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func InitProvider(w *zksync2.Wallet) (*zksync2.DefaultEthProvider, error) {
	ethRpc, err := rpc.Dial(rpcClientAddress)
	if err != nil {
		return nil, err
	}

	// and use it to create Ethereum Provider by Wallet
	ethereumProvider, err := w.CreateEthereumProvider(ethRpc)
	if err != nil {
		return nil, err
	}

	return ethereumProvider, nil
}
