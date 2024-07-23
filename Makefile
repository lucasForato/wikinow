install:
	templ generate && go install .
test:
	cd && cd Projects/wikinow/ && go build main.go && cd parser && go test -v
