FROM golang:1.15-alpine as builder

WORKDIR /build
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o passphrase-backend .

FROM scratch

COPY --from=builder /build/passphrase-backend /

CMD [ "/passphrase-backend" ]
