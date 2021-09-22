package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/registry"
	"gitlab.upay.dev/golang/kvstore/constant"
	"time"
)

type Postgresql struct {
	mCM   iface.ConfigMap
	mDSN  string
	mConn *pgx.Conn
}

func (p *Postgresql) Name() string {
	return Name
}

func (p *Postgresql) Version() string {
	return constant.Version
}

func (p *Postgresql) Category() string {
	return Category
}

func (p *Postgresql) ContractId() string {
	return ContractId
}

func (p *Postgresql) GetConfigMap() iface.ConfigMap {
	return p.mCM
}

func (p *Postgresql) New() iface.ICapability {
	return &Postgresql{}
}

func (p *Postgresql) SetConfigMap(cm iface.ConfigMap) error {
	p.mCM = cm
	p.mDSN = cm.String("db_dsn", "postgres://postgres:postgres@localhost:5432")
	return nil
}

func (p *Postgresql) Setup() error {
	conn, err := pgx.Connect(context.Background(), p.mDSN)
	if err != nil {
		return err
	}
	p.mConn = conn
	return nil
}

func (p *Postgresql) Get(ctx context.Context, key string, value interface{}) error {
	return nil
}

func (p *Postgresql) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (p *Postgresql) Delete(ctx context.Context, key string) error {
	return nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&Postgresql{})
}
