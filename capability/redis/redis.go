package redis

import (
	"context"
	"time"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/registry"
)

type KVRedis struct {
}

func (k *KVRedis) Name() string {
	return Name
}

func (k *KVRedis) Version() string {
	return Version
}

func (k *KVRedis) Category() string {
	return Category
}

func (k *KVRedis) ContractId() string {
	return ContractId
}

func (k *KVRedis) New() iface.ICapability {
	return &KVRedis{}
}

func (k *KVRedis) Get(ctx context.Context, key string, value interface{}) error {
	return nil
}

func (k *KVRedis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (k *KVRedis) Delete(ctx context.Context, key string) error {
	return nil
}

func init() {
	registry.NewCapabilityRegistry().RegisterCapability(ContractId, &KVRedis{})
}
