package main

import (
	_ "embed"

	"github.com/mkawserm/abesh/cmd"

	_ "github.com/mkawserm/httpserver2/capability/httpserver2"

	_ "github.com/amjadjibon/kvstore/capability/tikv"
	_ "github.com/amjadjibon/kvstore/encoding"
	_ "github.com/amjadjibon/kvstore/example/exTiKVSet"
)

//go:embed manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	cmd.Execute()
}
