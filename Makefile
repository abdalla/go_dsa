test:
	go test -v -count=1 ./...

run:
	go run --race cmd/main.go