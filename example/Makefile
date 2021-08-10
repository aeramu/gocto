generate:
	go mod init github.com/aeramu/gocto/example
	gocto generate
	mockery --all
	go mod tidy
mock:
	mockery --all --dir service
test:
	go test ./... --cover
