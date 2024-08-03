
install:
	templ generate && go install .

test:
	cd ~/Projects/wikinow/ && go test -v ./...

testq:
	cd ~/Projects/wikinow/ && go test ./...

