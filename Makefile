VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@git add VERSION
	@git commit -m "v$(VERSION)"
	@git tag -a "v$(VERSION)" -m "v$(VERSION)"
	@git push origin "v$(VERSION)"

test:
	go test -count=1 -race ./... -v

bench:
	go test -count=1 -race ./... -v -bench=. -benchtime=5s

update-module:
	go env -w GOPRIVATE=github.com/mkawserm
	go get -v github.com/mkawserm/abesh
	go get -u github.com/vmihailenco/msgpack/v5
	go get -u google.golang.org/protobuf
	go get -u github.com/go-redis/redis/v8
	go get -u github.com/jackc/pgx/v4
	go get -u github.com/google/flatbuffers/go

protoc:
	@protoc \
		-I=./proto \
			--go_opt=module=gitlab.upay.dev/golang/kvstore \
			--go_out=. \
			./proto/kvstore.proto

