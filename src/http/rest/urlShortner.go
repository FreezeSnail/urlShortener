package urlShortener

import (
	"fmt"
	"net/http"
)

type URLShortnerServer struct {
}

func (t URLShortnerServer) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer
	fmt.Print("Request Recieved\n")
	fmt.Printf("Request %v\n", r)
}
