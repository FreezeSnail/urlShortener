package main

import (
	"fmt"
	"net/http"

	api "github.com/FreezeSnail/urlShortener/src/http/rest"
)

func main() {
	s := api.URLShortnerServer{}
	h := api.Handler(s)

	fmt.Print("Hello world\n")
	http.ListenAndServe(":3000", h)
	fmt.Print("goodbye\n")
}
