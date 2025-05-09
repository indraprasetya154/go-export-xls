FROM golang:1.21.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/go-export-xls

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /out/go-export-xls /app/go-export-xls

CMD ["/app/go-export-xls"]
