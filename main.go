package main

import "fmt"

func main() {
	wallet, err := InitWallet()
	if err != nil {
		panic(err)
	}

	//provider, err := InitProvider(wallet)
	if err != nil {
		panic(err)
	}
	bal, err := wallet.GetBalance()
	if err != nil {
		panic(err)
	}
	fmt.Println(bal)
	//contract := DeployContract(wallet, provider)
	//fmt.Println(contract)
	TransferETH(wallet)

}
