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
)

type TiKV struct {
	mCM           model.ConfigMap
	mClient       *rawkv.Client
	mPDAddress    []string
	mEncodingName string
	mEncoding     encodingIface.IEncoding
}

func (t TiKV) Setup() error {
	var client, err = rawkv.NewClient(context.Background(), t.mPDAddress, config.DefaultConfig().Security)
	if err != nil {
		return err
	}
	t.mClient = client
	return nil
}

func (t TiKV) SetConfigMap(cm model.ConfigMap) error {
	t.mCM = cm
	t.mEncodingName = cm.String("encoding_name", "json")
	t.mPDAddress = cm.StringList("pd_address", ",", []string{})
	return nil
}

func (t TiKV) GetConfigMap() model.ConfigMap {
	return t.mCM
}

func (t TiKV) Name() string {
	return Name
}

func (t TiKV) Version() string {
	return Version
}

func (t TiKV) Category() string {
	return Category
}

func (t TiKV) ContractId() string {
	return ContractId
}

func (t TiKV) New() iface.ICapability {
	return &TiKV{}
}

func (t TiKV) Get(ctx context.Context, key string, value interface{}) error {
	valueBytes, err := t.mClient.Get(ctx, []byte(key))
	if err != nil {
		return err
	}
	return t.mEncoding.Unmarshal(valueBytes, value)
}

func (t TiKV) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	valueBytes, err := t.mEncoding.Marshal(value)
	if err != nil {
		return err
	}
	return t.mClient.PutWithTTL(ctx, []byte(key), valueBytes, uint64(ttl))
}

func (t TiKV) Delete(ctx context.Context, key string) error {
	return t.mClient.Delete(ctx, []byte(key))
}

func init() {
	registry.GlobalRegistry().AddCapability(&TiKV{})
}
