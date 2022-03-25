package utility

import (
	"github.com/mkawserm/abesh/iface"

	"github.com/amjadjibon/kvstore/capability/tikv"
)

// GetKVStoreTiKVCapability returns "golang:kvstore:tikv"
func GetKVStoreTiKVCapability(capabilityRegistry iface.ICapabilityRegistry) *tikv.TiKV {
	var r = capabilityRegistry.Capability(tikv.ContractId)
	if r == nil {
		return nil
	}

	var tiKV, ok = r.(*tikv.TiKV)
	if ok {
		return tiKV
	}

	return nil
}
