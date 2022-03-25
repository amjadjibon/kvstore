package tikv

import (
	"context"
	"time"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
	"github.com/tikv/client-go/v2/rawkv"
	"github.com/tikv/client-go/v2/txnkv"

	"github.com/amjadjibon/kvstore/errors"
	kvStoreIFace "github.com/amjadjibon/kvstore/iface"
	kvStoreRegistry "github.com/amjadjibon/kvstore/registry"
)

type TiKV struct {
	mCM           model.ConfigMap
	mEncoding     kvStoreIFace.IEncoding
	mEncodingName string
	mClient       *txnkv.Client
	mRawKVClient  *rawkv.Client
	mPDAddress    []string
}

func (t *TiKV) Name() string {
	return Name
}

func (t *TiKV) Version() string {
	return Version
}

func (t *TiKV) Category() string {
	return Category
}

func (t *TiKV) ContractId() string {
	return ContractId
}

func (t *TiKV) New() iface.ICapability {
	return &TiKV{}
}

func (t *TiKV) Setup() error {
	t.mEncoding = kvStoreRegistry.EncodingRegistry().GetEncoding(t.mEncodingName)
	if t.mEncoding == nil {
		return errors.ErrEncodingNotFound
	}

	var client, err = txnkv.NewClient(t.mPDAddress)
	if err != nil {
		return err
	}
	t.mClient = client

	return nil
}

func (t *TiKV) SetConfigMap(cm model.ConfigMap) error {
	t.mCM = cm
	t.mEncodingName = cm.String("encoding_name", "json")
	t.mPDAddress = cm.StringList("pd_address", ",", []string{"localhost:2379"})
	return nil
}

func (t *TiKV) GetConfigMap() model.ConfigMap {
	return t.mCM
}

func (t *TiKV) Get(ctx context.Context, key string, value interface{}) error {
	tx, err := t.mClient.Begin()
	if err != nil {
		return err
	}

	data, err := tx.Get(ctx, []byte(key))
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return t.mEncoding.Unmarshal(data, value)
}

func (t *TiKV) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	v, err := t.mEncoding.Marshal(value)
	if err != nil {
		return err
	}

	tx, err := t.mClient.Begin()
	if err != nil {
		return err
	}

	err = tx.Set([]byte(key), v)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (t *TiKV) Delete(ctx context.Context, key string) error {
	tx, err := t.mClient.Begin()
	if err != nil {
		return err
	}

	err = tx.Delete([]byte(key))
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func init() {
	registry.GlobalRegistry().AddCapability(&TiKV{})
}
