build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go mod tidy
	go run .

icon:
	go install github.com/akavel/rsrc@latest
	rsrc -ico go.ico

compile: icon
	go mod tidy
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o bin/rasterize-linux-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/rasterize-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/rasterize-macos-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/rasterize-windows-amd64.exe .

