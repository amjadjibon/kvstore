package rawkv

import (
	"context"
	"time"

	encodingIface "github.com/amjadjibon/encoding/iface"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"

	"github.com/amjadjibon/kvstore/tikv/constant"
)

type TiKVRawKV struct {
	mCM           model.ConfigMap
	mClient       *rawkv.Client
	mPDAddress    []string
	mEncodingName string
	mEncoding     encodingIface.IEncoding
}

func (t *TiKVRawKV) Name() string {
	return NameRawKV
}

func (t *TiKVRawKV) Version() string {
	return constant.Version
}

func (t *TiKVRawKV) Category() string {
	return Category
}

func (t *TiKVRawKV) ContractId() string {
	return ContractIdRawKV
}

func (t *TiKVRawKV) New() iface.ICapability {
	return &TiKVRawKV{}
}

func (t *TiKVRawKV) Setup() error {
	var client, err = rawkv.NewClient(context.Background(), t.mPDAddress, config.DefaultConfig().Security)
	if err != nil {
		return err
	}
	t.mClient = client
	return nil
}

func (t *TiKVRawKV) SetConfigMap(cm model.ConfigMap) error {
	t.mCM = cm
	t.mEncodingName = cm.String("encoding_name", "json")
	t.mPDAddress = cm.StringList("pd_address", ",", []string{})
	return nil
}

func (t *TiKVRawKV) GetConfigMap() model.ConfigMap {
	return t.mCM
}

func (t *TiKVRawKV) Get(ctx context.Context, key string, value interface{}) error {
	valueBytes, err := t.mClient.Get(ctx, []byte(key))
	if err != nil {
		return err
	}
	return t.mEncoding.Unmarshal(valueBytes, value)
}

func (t *TiKVRawKV) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	valueBytes, err := t.mEncoding.Marshal(value)
	if err != nil {
		return err
	}
	return t.mClient.PutWithTTL(ctx, []byte(key), valueBytes, uint64(ttl))
}

func (t *TiKVRawKV) Delete(ctx context.Context, key string) error {
	return t.mClient.Delete(ctx, []byte(key))
}

func init() {
	registry.GlobalRegistry().AddCapability(&TiKVRawKV{})
}
