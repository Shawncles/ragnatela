// Command eval is a chromedp example demonstrating how to evaluate javascript
// and retrieve the result.
package main
// FIXME render sia explorer
import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	"fmt"

)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	fmt.Println("create instance\n")
	//run, err := runner.New(runner.Flag("headless", true),
	//	  				   runner.Flag("no-sandbox", true))
	//run, err := runner.New(runner.Flag("headless", true))
	//c, err := chromedp.New(ctxt, chromedp.WithRunner(run))
	c, err := chromedp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var res []string
	fmt.Printf("starting running\n")

	err = c.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(`https://explorer.siahub.info/hash/faec097bfb94939dc0ade93586a307b81a7a9be0b1f25e37ead967f6249ed236c9e9209f3011`),
		chromedp.WaitVisible(`.row`, chromedp.ByID),
		chromedp.Evaluate(`Object.keys(window);`, &res),
	})
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("window object keys: %v", res)
}