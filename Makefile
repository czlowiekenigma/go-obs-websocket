all:
	go build

check:
	go vet
	go test -coverprofile=cover.out
	golint
