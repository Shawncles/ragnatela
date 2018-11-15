package api

import
(
	"net/http"
	"log"
)

func SetupServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/address", BalanceHandler)
	mux.HandleFunc("/transaction", TansferHandler)
	log.Println("Listening...")
	http.ListenAndServe(":2333", mux)
}
