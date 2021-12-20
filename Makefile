# To enable hot reload of our server
build: 
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run