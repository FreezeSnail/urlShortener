package main

import (
	"net/http"

	api "github.com/FreezeSnail/urlShortener/http/rest"
)

func main() {
	s := api.URLShortnerServer{}
	h := api.Handler(s)

	http.ListenAndServe(":3000", h)
}
