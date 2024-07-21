build:
	go build -o bin/to_do_list cmd/*/main.go
run: build
	./bin/to_do_list