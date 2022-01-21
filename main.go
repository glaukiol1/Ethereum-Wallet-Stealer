package main

import (
	"fmt"
	"math/big"

	"github.com/glaukiol1/Ethereum-Wallet-Stealer/src"
)

var addresses []string
var privateKeys []string

func main() {
	for i := 0; true; i++ {
		pb, pk := src.Keygen()
		_pb := src.PubkeyToAddress(pb)
		_pk := src.PrivateKeyToHex(pk)
		addresses = append(addresses, _pb)
		privateKeys = append(privateKeys, _pk)
		if i == 1000 {
			i = 0
			for j := 0; j < len(addresses); j++ {
				adr := addresses[j]
				pkr := privateKeys[j]
				if src.HasBalance(adr).Cmp(big.NewFloat(0)) == 1 {
					fmt.Println("HIT! Address: " + adr + " Seed: " + pkr)
				} else {
					fmt.Println("Missed!")
				}
			}
		}
	}
}
