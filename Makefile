VERSION = "0.0.1"
change-version:
	@echo $(version)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(version)\"">constant/version.go
	@git add VERSION
	@git add constant/version.go
	@git commit -m "v$(version)"
	@git tag -a "v$(version)" -m "v$(version)"
	@git push origin "v$(version)"

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

