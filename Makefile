build: build-ui build-api

build-api:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o droplify *.go

build-ui:
	cd ui && npm run build && cd .. && ./files2go