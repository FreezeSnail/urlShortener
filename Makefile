GOARCH = amd64

all: gen build

gen:
	oapi-codegen -generate types -o domain/openapi_types.gen.go -package types openapi/urlShortener.yml
	oapi-codegen -generate chi-server -o http/rest/openapi_server.gen.go -package api openapi/urlShortener.yml

build:
	go build -o urlShortner cmd/api/main.go 