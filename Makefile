VERSION = "0.0.1"
change-tikv-version:
	@echo "package tikv\n\n// Version constant of tikv\nconst Version = \"$(VERSION)\"">tikv/capability/rawkv/version.go
	@git add tikv/capability/rawkv/version.go
	@git commit -m "tikv/v$(VERSION)"
	@git tag -a "tikv/v$(VERSION)" -m "tikv/v$(VERSION)"
	@git push origin
	@git push origin "tikv/v$(VERSION)"

test:
	go test -count=1 -race ./... -v

bench:
	go test -count=1 -race ./... -v -bench=. -benchtime=5s

update-module:
	go env -w GOPRIVATE=github.com/mkawserm
	go get -v github.com/mkawserm/abesh
