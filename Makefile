VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION

test:
	go test -count=1 -race ./... -v

update-module:
	go env -w GOPRIVATE=github.com/mkawserm,gitlab.upay.dev/golang
	go get -v github.com/mkawserm/abesh
