package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

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

type s1 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
	} `json:"result"`
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func BalanceList(addressList []string, pkList []string) {
	list := chunkSlice(addressList, 20)
	for _, lst := range list {
		resp, err := http.Get(
			"https://api.etherscan.io/api?module=account&action=balancemulti&address=" + strings.Join(lst, ",") + "&tag=latest&apikey=HVIFHDEQKUUCMPB7FB6CHGBC7CPG6ZIWX6")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var res s1
		if resp.StatusCode != 200 {
			panic("Non-200 Status Code")
		}
		err = json.Unmarshal(body, &res)
		if err != nil {
			if strings.Index(string(body), "Max rate limit reached") != -1 {
				println("Rate limited")
				continue
			} else {
				panic(err)
			}
		}
		for i := 0; i < len(res.Result); i++ {
			current := res.Result[i]
			n, _ := strconv.Atoi(current.Balance)
			bal := weiToEther(big.NewInt(int64(n)))
			if bal.Cmp(big.NewFloat(0)) == 1 {
				fmt.Println("Hit! \n\r Public Key: " + current.Account + "\n\r Private Key: " + pkList[i] + "\n\r Find this data in SUCCESS.txt")
				f, _ := os.Create("./SUCCESS.txt")
				f.WriteString("Address: " + current.Account + "\nPrivate Key: " + pkList[i])
			}
		}
	}
}
