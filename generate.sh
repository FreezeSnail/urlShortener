oapi-codegen -generate types -o openapi_types.gen.go -package main urlShortner.yml
oapi-codegen -generate chi-server -o openapi_server.gen.go -package main urlShortner.yml