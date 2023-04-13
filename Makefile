GOARCH = amd64

all: db urlshortener
dev: gen build

gen:
	oapi-codegen -generate chi-server -o src/http/rest/urlShortner.gen.go -package urlShortener src/openapi/urlShortener.yml
	oapi-codegen -generate types -o src/http/rest/urlShortener_types.gen.go -package urlShortener src/openapi/urlShortener.yml

build:
	go build -o urlShortner src/cmd/api/main.go 

db:
	docker build -t urlshortener/sqlite src/db/
urlshortener:
	docker build -t urlshortener/urlshortener .