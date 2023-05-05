package urlShortener

import (
	"encoding/json"
	"fmt"
	"net/http"

	shortener "github.com/FreezeSnail/urlShortener/src/cmd/shortener"
	sqlite "github.com/FreezeSnail/urlShortener/src/db"
	domain "github.com/FreezeSnail/urlShortener/src/domain"
)

type URLShortnerServer struct {
	DB *sqlite.SQLite
}

func (t URLShortnerServer) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var body domain.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}
	// our logic to store the todo into a persistent layer
	fmt.Print("Request Recieved\n")

	//check if url is in db -> should be in middleware?

	//shorten url
	short, err := shortener.Shorten(body.Url)
	if err != nil {
		panic(err)
	}

	//store url in db
	resp, err := t.DB.AddUrl(body.Url, short)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}
	//return shortened url
	writeJSON(w, resp)

}

func (t URLShortnerServer) getShortnedUrl(w http.ResponseWriter, r *http.Request) {
	var body domain.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}
	// our logic to store the todo into a persistent layer
	fmt.Print("Request Recieved\n")

	//check if url is in db -> should be in middleware?

	//shorten url
	short, err := shortener.Shorten(body.Url)
	if err != nil {
		panic(err)
	}

	//store url in db
	resp, err := t.DB.AddUrl(body.Url, short)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}
	//return shortened url
	writeJSON(w, resp)

}

func writeError(w http.ResponseWriter, code int, err error) {
	type response struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{Error: err.Error()})
}

// writeJSON is a helper function to write a JSON response.
func writeJSON(w http.ResponseWriter, b any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
