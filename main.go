package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

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
		if i == 1000000 {
			break
		}
	}
	eachFile := len(addresses) / 3
	pb1, _ := os.Create("./public1.out")
	pk1, _ := os.Create("./private1.out")

	pb1.WriteString(strings.Join(addresses[0:eachFile], "\n"))
	pk1.WriteString(strings.Join(privateKeys[0:eachFile], "\n"))

	pb2, _ := os.Create("./public2.out")
	pk2, _ := os.Create("./private2.out")

	pb2.WriteString(strings.Join(addresses[0:eachFile], "\n"))
	pk2.WriteString(strings.Join(privateKeys[0:eachFile], "\n"))

	pb3, _ := os.Create("./public3.out")
	pk3, _ := os.Create("./private3.out")

	pb3.WriteString(strings.Join(addresses[0:eachFile], "\n"))
	pk3.WriteString(strings.Join(privateKeys[0:eachFile], "\n"))

	for j := 0; j < len(addresses); j++ {
		adr := addresses[j]
		pkr := privateKeys[j]
		if src.HasBalance(adr).Cmp(big.NewFloat(0)) == 1 {
			fmt.Println("HIT! Address: " + adr + " Seed: " + pkr)
		} else {
			fmt.Println("Missed! " + adr)
		}
	}
}
