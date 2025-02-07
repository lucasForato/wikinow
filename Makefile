templ-generate:
	templ generate

tailwind-cli:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.14/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

dev:
	make templ-generate 
	go install . 
	wikinow start

test:
	go test -v ./...

test-race:
	go test -race -v ./...

test-cover:
	go test -v -cover ./...

install:
	make templ-generate
	go install .

build:
	go build -o wikinow main.go
