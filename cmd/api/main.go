package main

import (
	"fmt"
	"net/http"

	api "github.com/FreezeSnail/urlShortener/http/rest"
)

type URLShortnerServer struct {
}

func (t URLShortnerServer) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer
	fmt.Printf("Request %v", r)
}

func main() {
	s := URLShortnerServer{}
	h := api.Handler(s)

	http.ListenAndServe(":3000", h)
}
