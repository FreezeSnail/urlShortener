package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/FreezeSnail/urlShortener/src/db"
	api "github.com/FreezeSnail/urlShortener/src/http"
	domain "github.com/FreezeSnail/urlShortener/src/http/domain"
	middles "github.com/FreezeSnail/urlShortener/src/http/middles"
	"github.com/go-chi/chi"
	"golang.org/x/exp/slog"
)

func main() {

	db, err := db.NewSQLite("urlshortener.db")
	if err != nil {
		db.Close()
		panic(err)
	}
	defer db.Close()
	textHandler := slog.NewTextHandler(os.Stdout)
	logger := slog.New(textHandler)

	s := api.URLShortnerServer{
		DB:  db,
		Log: logger,
	}

	m := middles.Middleware{
		DB:  db,
		Log: logger,
	}
	r := chi.NewRouter()
	r.Mount("/", domain.Handler(&s, domain.WithMiddleware("validateAPIKey", m.ValidateAPIKey())))
	//h := api.Handler(s)

	//TODO(SNAIL) Add background process to check token ttl

	fmt.Print("Hello world\n")
	http.ListenAndServe(":3000", r)
	fmt.Print("goodbye\n")
}
