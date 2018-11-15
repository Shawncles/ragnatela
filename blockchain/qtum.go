package blockchain

import (
	"fmt"
	"regexp"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"strconv"
)

type Qtum struct {
	Tokens []Token
}

type Token struct{
	Address 		string 	`json:"address"`
	Name 			string	`json:"name"`
	Symbol 			string	`json:"symbol"`
	ContractAddress string	`json:"contract_address"`
	Balance 		int64	`json:"balance"`
}

func (qtum *Qtum) ListBalances(address string) (err error)  {
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		exp := regexp.MustCompile(`[b|B]alance\:\s?\"?([0-9]+?)\"`)
		target := exp.FindAllStringSubmatch(string(r.Body), -1)

		qtum.Tokens = append(qtum.Tokens, Token{address, "Loopring Qtum Token", "QTUM", "2eb2a66afd4e465fb06d8b71f30fb1b93e18788d", -1})
		qtum.Tokens[0].Balance, err = strconv.ParseInt(target[0][1], 10, 64)

		fmt.Printf("qtum.Tokens[0].Address is: %s\n", qtum.Tokens[0])
		fmt.Printf("qtum.Tokens[0].Balance is: %d\n", qtum.Tokens[0].Balance)
		if len(target) != 1 {
			//exp := regexp.MustCompile(`qrc20TokenBalances\:\[([\w|\:|\"|\,|\{|\}|\s]+?)\]`)
			exp := regexp.MustCompile(`address\:\"([\w]+)\"\,name\:\"([\w|\s]+)\"\,symbol\:\"(\w+)\"\,[\w|\:|\"]+\,[\w|\:|\"]+\,balance\:\"(\d+)\"`)
			target := exp.FindAllStringSubmatch(string(r.Body), -1)
			var tk Token
			//fmt.Println("tk is ..........")
			for i:=0; i < len(target); i++ {
				token := target[i]
				//fmt.Printf("tk is %s\n", target[i])
				tk.ContractAddress = token[1]
				tk.Name = token[2]
				tk.Symbol = token[3]
				tk.Balance, err = strconv.ParseInt(token[4], 10, 64)
				tk.Address = address
				qtum.Tokens = append(qtum.Tokens, tk)
			}
		}
	})
	c.Visit("https://qtum.info/address/" + address)
	c.Wait()
	return nil
}