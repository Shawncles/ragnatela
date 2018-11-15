package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../blockchain"
	"strings"
)

type BalanceQuery struct{
	Chain		string `json:"chain"`
	Address		string `json:"address"`
}

type Balance struct {
	Chain 			string `json:"chain"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contract_address"`
	Balance         int64  `json:"balance"`
}


var BalanceHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Balance query received request")
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		fmt.Printf("r.Body is %s\n", string(body))
		var q BalanceQuery
		err = json.Unmarshal(body, &q)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		fmt.Println(q)
		balanceInfo, err := q.GetAddressBalance()
		fmt.Printf("balanceInfo is %s\n", balanceInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
		balanceInfoJSON, err := json.Marshal(balanceInfo)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(balanceInfoJSON)
		w.WriteHeader(200)
		fmt.Println("server: OK")
		return
	default:
		http.Error(w, "Unsupport method error", http.StatusInternalServerError)
		return
	}
}

type BC struct {
	blockchain.BlockChain
}

func (q * BalanceQuery) GetAddressBalance() (balance []Balance, err error) {
	switch strings.ToLower(q.Chain) {
		case "qtum":
			qtum := blockchain.Qtum{}
			qtum.ListBalances(q.Address)
			var balance = make([]Balance, len(qtum.Tokens))
			for i := 0; i < len(qtum.Tokens); i++ {
				token := qtum.Tokens[i]
				balance[i].Chain 	= "qtum"
				balance[i].Address 	= token.Address
				balance[i].Name 	= token.Name
				balance[i].Symbol 	= token.Symbol
				balance[i].ContractAddress = token.ContractAddress
				balance[i].Balance = token.Balance
			}
			fmt.Println(balance)
			fmt.Println(len(qtum.Tokens))
			return balance, nil
		case "siacore":
			siacore := blockchain.SiaCore{}
			siacore.ListBalances(q.Address)
			var balance = make([]Balance, 1)
			balance[0].Address = siacore.Address
			balance[0].Balance = siacore.Balance
			return balance, nil
	}
	return
}


