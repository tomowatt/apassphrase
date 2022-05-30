ARG GO_BUILD_ENV

FROM golang:1.18-alpine as builder

WORKDIR /build
COPY . .
RUN $GO_BUILD_ENV go build -o passphrase-backend .

FROM alpine:3.15

WORKDIR /app/
COPY --from=builder /build/passphrase-backend .
RUN adduser -S passphrase
USER passphrase
CMD ./passphrase-backend
