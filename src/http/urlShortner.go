package urlShortener

import (
	"encoding/json"
	"net/http"

	shortener "github.com/FreezeSnail/urlShortener/src/cmd/shortener"
	sqlite "github.com/FreezeSnail/urlShortener/src/db"
	domain "github.com/FreezeSnail/urlShortener/src/http/domain"
	bcrypt "golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

type URLShortnerServer struct {
	DB  *sqlite.SQLite
	Log *slog.Logger
}

func (t URLShortnerServer) ShortenURL(w http.ResponseWriter, r *http.Request) *domain.Response {
	t.Log.Info("request recieved to shorten url")
	var body domain.ShortenURLRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}
	// our logic to store the todo into a persistent layer
	//check if url is in db -> should be in middleware?

	//shorten url
	short, err := shortener.Shorten(body.URL)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	//store url in db
	resp, err := t.DB.AddUrl(r.Context(), body.URL, short)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}
	//return shortened url
	return domain.ShortenURLJSON201Response(*resp)

}

func (t URLShortnerServer) GetLongURLFromShort(w http.ResponseWriter, r *http.Request, url string) *domain.Response {
	t.Log.Info("request recieved to get short url")

	resp, err := t.DB.GetLongUrl(r.Context(), url)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return nil
	}
	//return shortened url
	return domain.GetLongURLFromShortJSON200Response(*resp)

}

func (t URLShortnerServer) PostSignUp(w http.ResponseWriter, r *http.Request) *domain.Response {
	t.Log.Info("request recieved to sign up")
	var body domain.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 1)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	// get apiKey generation
	apiKey := "key"
	err = t.DB.CreateUser(r.Context(), body.Username, string(hashedPass), apiKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	resp := &domain.SignInResponse{
		Apikey: &apiKey,
	}
	return domain.PostSignInJSON201Response(*resp)
}

func (t URLShortnerServer) PostSignIn(w http.ResponseWriter, r *http.Request) *domain.Response {
	t.Log.Info("request recieved to sign up")
	var body domain.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	hashedPass, err := t.DB.GetHashedPassword(r.Context(), body.Username)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(body.Password))
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return nil
	}

	apiKey, err := t.DB.GetAPIKey(r.Context(), body.Username, hashedPass)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return nil
	}

	resp := &domain.SignInResponse{
		Apikey: &apiKey,
	}

	return domain.PostSignInJSON201Response(*resp)

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
