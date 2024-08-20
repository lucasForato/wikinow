install:
	templ generate && go install .

test:
	cd ~/Projects/wikinow/ && go test -v ./...

testq:
	cd ~/Projects/wikinow/ && go test ./...

tailwind:
	npx tailwindcss -i static/css/input.css -o static/css/style.css
