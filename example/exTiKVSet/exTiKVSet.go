package exTiKVSet

import (
	"context"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/kvstore/example"
	"github.com/amjadjibon/kvstore/utility"
)

type ExTiKVSet struct {
	mValues  model.ConfigMap
	mKVStore iface.IKVStore
}

func (e *ExTiKVSet) Name() string {
	return "abesh_kvstore_tikv_set_example"
}

func (e *ExTiKVSet) Version() string {
	return "0.0.1"
}

func (e *ExTiKVSet) Category() string {
	return string(constant.CategoryService)
}

func (e *ExTiKVSet) ContractId() string {
	return "abesh:kvstore:exTiKVSet"
}

func (e *ExTiKVSet) GetConfigMap() model.ConfigMap {
	return e.mValues
}

func (e *ExTiKVSet) Setup() error {
	return nil
}

func (e *ExTiKVSet) SetConfigMap(values model.ConfigMap) error {
	e.mValues = values
	return nil
}

func (e *ExTiKVSet) New() iface.ICapability {
	return &ExTiKVSet{}
}

func (e *ExTiKVSet) SetCapabilityRegistry(capabilityRegistry iface.ICapabilityRegistry) error {
	e.mKVStore = utility.GetKVStoreTiKVCapability(capabilityRegistry)
	return nil
}

func (e *ExTiKVSet) Serve(ctx context.Context, input *model.Event) (*model.Event, error) {
	return example.ExPostServe(ctx, input, e.mKVStore, e.ContractId())
}

func init() {
	registry.GlobalRegistry().AddCapability(&ExTiKVSet{})
}
