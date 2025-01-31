FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 go build -o ultradex .

FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app/

COPY --from=build /app/ultradex .

COPY --from=build /app/database/ /app/database/
COPY --from=build /app/web/ /app/web/
COPY --from=build /app/fly.toml /app/fly.toml

EXPOSE 8080/tcp

CMD ["./ultradex"]