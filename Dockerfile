FROM golang:1.15-alpine as builder

WORKDIR /build
COPY . . 
RUN GOPATH=$PWD CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o emojiphrase-backend .

FROM alpine:3.12

WORKDIR /app/
COPY --from=builder /build/emojiphrase-backend .
RUN adduser -S passphrase
USER passphrase
CMD ./emojiphrase-backend
