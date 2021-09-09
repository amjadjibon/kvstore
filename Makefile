VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(VERSION)\"">constant/version.go

test:
	go test -count=1 -race ./... -v

bench:
	go test -count=1 -race ./... -v -bench=. -benchtime=5s

update-module:
	go env -w GOPRIVATE=github.com/mkawserm,gitlab.upay.dev/golang
	go get -v github.com/mkawserm/abesh
	go get -u github.com/vmihailenco/msgpack/v5
	go get -u google.golang.org/protobuf
	go get -u github.com/go-redis/redis/v8

update-protoc:
	@protoc \
		-I=./proto \
			--go_opt=module=gitlab.upay.dev/golang/kvstore \
			--go_out=. \
			./proto/kvstore.proto

