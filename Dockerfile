FROM golang:1.20 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o lamoda cmd/lamoda/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=build app/lamoda .
COPY scheme ./scheme

EXPOSE 9000 8000

CMD ["./lamoda"]
