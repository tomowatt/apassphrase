FROM golang:1.15-alpine as builder

WORKDIR /build
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o passphrase-backend .

FROM alpine:3.14

WORKDIR /app/
COPY --from=builder /build/passphrase-backend .
RUN adduser -S passphrase
USER passphrase
CMD ./passphrase-backend
