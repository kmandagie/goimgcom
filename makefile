test:
	go test  ./...

coverage:
	go test -coverprofile cp.out ./...
	go tool cover -html=cp.out