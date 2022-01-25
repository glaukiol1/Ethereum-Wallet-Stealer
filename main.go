package main

import (
	"github.com/glaukiol1/Ethereum-Wallet-Stealer/src"
)

var addresses []string
var privateKeys []string

func main() {
	hashes := 500000 // five hundred thousand 500,000
	for i := 0; true; i++ {
		pb, pk := src.Keygen()
		_pb := src.PubkeyToAddress(pb)
		_pk := src.PrivateKeyToHex(pk)
		addresses = append(addresses, _pb)
		privateKeys = append(privateKeys, _pk)
		if i == hashes {
			break
		}
	}
	src.Run(addresses, privateKeys)
}
