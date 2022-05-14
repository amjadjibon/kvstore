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
