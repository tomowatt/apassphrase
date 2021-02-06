GO_BUILD_ENV := CGO_ENABLED=0 GOARCH=amd64 GOPATH=$(PWD)

build:
	$(GO_BUILD_ENV) go build -o bin/emojiphrase-backend -v .

docker-build:
	docker build -t local/passphrase-backend .

docker-run:
	docker run -d -e PORT=8080 -p 8080:8080 local/passphrase-backend

clean:
	rm -rf bin/emojiphrase-backend
	docker rmi -f local/passphrase-backend
