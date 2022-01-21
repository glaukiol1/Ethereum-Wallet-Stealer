package src

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keygen() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	privateKey, err := crypto.GenerateKey()
	publicKey := privateKey.Public()
	if err != nil {
		log.Fatal(err)
	}
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return publicKeyECDSA, privateKey
}

func PubkeyToAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]
}
