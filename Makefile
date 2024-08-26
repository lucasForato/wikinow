tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css

templ-generate:
	templ generate

dev:
	make templ-generate 
	make tailwind-build 
	go install . 
	wikinow start

test:
	  go test -race -v -timeout 30s ./...

install:
	make templ-generate
	make tailwind-build
	go install .
