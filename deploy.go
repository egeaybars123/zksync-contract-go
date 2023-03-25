package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go"
)

func DeployContract(w *zksync2.Wallet, p *zksync2.DefaultEthProvider) string {
	hash, err := w.Deploy(bytecode, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction Hash: ", hash)

	address, err := zksync2.ComputeL2Create2Address(
		common.HexToAddress(GetEnv("DEPLOYER_L2_ADDRESS")),
		bytecode,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Contract deployed to: ", address.String())

	return address.String()
}
