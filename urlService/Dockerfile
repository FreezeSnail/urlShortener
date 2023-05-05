FROM golang:1.20

WORKDIR /app
COPY . . 

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN apk update && apk upgrade
RUN apk add --no-cache sqlite

RUN go mod download
RUN ls
RUN make gen && make build
ENTRYPOINT [ "/app/urlShortner" ]