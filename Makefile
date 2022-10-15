PROJECT:=vAdmin

.PHONY: build
build:
	CGO_ENABLED=0 go build -o vAdmin main.go
build-sqlite:
	go build -tags sqlite3 -o vAdmin main.go
#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t vAdmin:latest
