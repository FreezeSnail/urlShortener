package urlShortener

import (
	"fmt"
	"net/http"
)

//go:generate oapi-codegen -generate chi-server -o urlShortner.gen.go -package urlShortener ../../openapi/urlShortener.yml
//go:generate oapi-codegen -generate types -o urlShortener_types.gen.go -package urlShortener ../../openapi/urlShortener.yml

type URLShortnerServer struct {
}

func (t URLShortnerServer) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer
	fmt.Printf("Request %v", r)
}
