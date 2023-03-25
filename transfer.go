package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go"
)

func TransferETH(w *zksync2.Wallet) {
	hash, err := w.Transfer(
		common.HexToAddress(GetEnv("TO_TEST")),
		big.NewInt(1000000000000),
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx hash", hash)
}
