all:
	build test

run: 
	build run

build:
	go build -o dist/server cmd/server/main.go

test:
	go test ./...

start:
	./dist/server

clean:
	rm -fr ./dist/*