generate:
	go mod init github.com/aeramu/example
	gocto generate
	mockery --all
	go mod tidy
mock:
	mockery --all
test:
	go test ./...
