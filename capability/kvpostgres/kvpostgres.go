package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
	"gitlab.upay.dev/golang/kvstore/constant"
	"time"
)

type KVPostgres struct {
	mCM   model.ConfigMap
	mDSN  string
	mConn *pgx.Conn
	mPool *pgxpool.Pool
}

func (k *KVPostgres) Name() string {
	return constant.NameKVPostgres
}

func (k *KVPostgres) Version() string {
	return constant.Version
}

func (k *KVPostgres) Category() string {
	return constant.Category
}

func (k *KVPostgres) ContractId() string {
	return constant.ContractIdKVPostgres
}

func (k *KVPostgres) GetConfigMap() model.ConfigMap {
	return k.mCM
}

func (k *KVPostgres) New() iface.ICapability {
	return &KVPostgres{}
}

func (k *KVPostgres) SetConfigMap(cm model.ConfigMap) error {
	k.mCM = cm
	k.mDSN = cm.String("db_dsn", "")
	return nil
}

func (k *KVPostgres) Setup() error {
	conn, err := pgx.Connect(context.Background(), k.mDSN)
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), k.mDSN)

	if err != nil {
		return err
	}

	k.mConn = conn
	k.mPool = pool
	return nil
}

func (k *KVPostgres) Get(ctx context.Context, key string, value interface{}) error {
	err := k.mPool.QueryRow(ctx, "select key, value from kvstore where key = ?", key).Scan(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (k *KVPostgres) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (k *KVPostgres) Delete(ctx context.Context, key string) error {
	return nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&KVPostgres{})
}
