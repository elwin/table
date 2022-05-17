test:
	go test ./...

install:
	go install

release:
	goreleaser release
