build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

icon:
	go install github.com/akavel/rsrc@latest
	rsrc -ico go.ico

run:
	go mod tidy
	go run .

all: icon build
