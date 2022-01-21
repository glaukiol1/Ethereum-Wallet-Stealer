package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/params"
)

type s struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func HasBalance(address string) *big.Float {
	resp, err := http.Get(
		"https://api.etherscan.io/api?module=account&action=balance&address=" + address + "&tag=latest&apikey=HVIFHDEQKUUCMPB7FB6CHGBC7CPG6ZIWX6")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res s
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}
	intn, _ := strconv.Atoi(res.Result)
	return weiToEther(big.NewInt(int64(intn)))
}
