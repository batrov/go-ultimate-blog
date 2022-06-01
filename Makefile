install-air:
	go install github.com/cosmtrek/air@latest
	
build: 
	go build

run: 
	go run app.go

test:
	go test ./...

dev:
	air