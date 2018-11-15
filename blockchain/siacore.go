package blockchain

import (
	"github.com/gocolly/colly"
	"fmt"
)

type SiaCore struct {
	Address string 	`json:"address"`
	Balance int64	`json:"balance"`
}

var bodyData []byte
func (sc *SiaCore) ListBalances(address string) (err error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.MaxDepth(10),
		colly.Async(true),
		//colly.Debugger(&debug.LogDebugger{}),
	)
	//c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {

		//bodyData := r.Body
		fmt.Printf("colly.Response is %s\n", string(r.Body))
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	fmt.Println("calling onhtml ......... ")
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("using onhtml method.......")
		//e.DOM.rend
		//e = e.DOM.Find("Balance")
		//if e.Index == 2 {
		//	fmt.Printf("onhtml is %s\n", e.Text)
		//}
	})

	c.Visit("https://explorer.siahub.info/hash/" + address)
	//"https://explorer.siahub.info/hash/faec097bfb94939dc0ade93586a307b81a7a9be0b1f25e37ead967f6249ed236c9e9209f3011"
	return nil
}