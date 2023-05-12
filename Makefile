GOARCH = amd64

all: db urlshortener
dev: gen build

gen:
	goapi-gen -generate server -o src/http/domain/urlShortner.gen.go -package domain src/openapi/urlShortener.yml
	goapi-gen -generate types -o src/http/domain/urlShortener_types.gen.go -package domain src/openapi/urlShortener.yml
	go generate src/db



build:
	go build -o urlShortener ./src/cmd/api/main.go 

urlshortener:
	docker build -t urlshortener/urlshortener .