GOARCH = amd64

all: db urlshortener
dev: gen build

gen:
	oapi-codegen -generate chi-server -o src/http/rest/urlShortner.gen.go -package urlShortener src/openapi/urlShortener.yml
	oapi-codegen -generate types -o src/domain/urlShortener_types.gen.go -package domain src/openapi/urlShortener.yml
	go generate
	


build:
	go build -o urlShortner src/cmd/api/main.go 

urlshortener:
	docker build -t urlshortener/urlshortener .