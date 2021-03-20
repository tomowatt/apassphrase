GO_BUILD_ENV := CGO_ENABLED=0 GOARCH=amd64
LOCAL_ENV := PORT=8080

build:
	$(GO_BUILD_ENV) go build -o bin/passphrase-backend -v .

local-run: build
	$(LOCAL_ENV) bin/passphrase-backend

docker-build:
	docker build -t local/passphrase-backend . --no-cache

docker-run: docker-build
	docker run -d -e $(LOCAL_ENV) -p 8080:8080 --name local-passphrase-backend local/passphrase-backend

clean:
	rm -rf bin/passphrase-backend
	docker rm -f local-passphrase-backend
	docker rmi -f local/passphrase-backend
