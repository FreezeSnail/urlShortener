GOARCH = amd64

all: gen build

gen:
	oapi-codegen -generate types -o pkg/types/openapi_types.gen.go -package types openapi/urlShortner.yml
	oapi-codegen -generate chi-server -o pkg/api/openapi_server.gen.go -package api openapi/urlShortner.yml

build:
	go build -o urlShortner main.go 