GOARCH = amd64

all: gen build

gen:
	go generate
	oapi-codegen -generate types -o domain/urlShortener_types.gen.go -package types openapi/urlShortener.yml

build:
	go build -o urlShortner cmd/api/main.go 