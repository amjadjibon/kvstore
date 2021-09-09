package kvredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/registry"
	"gitlab.upay.dev/golang/kvstore/constant"
	"time"
)

type Redis struct {
	/* REDIS CONNECTION OPTIONS*/
	mAddress            string
	mUserName           string
	mPassword           string
	mDB                 int
	mNetwork            string
	mMaxRetries         int
	mMinRetryBackoff    time.Duration
	mMaxRetryBackoff    time.Duration
	mDialTimeout        time.Duration
	mReadTimeout        time.Duration
	mWriteTimeout       time.Duration
	mPoolFIFO           bool
	mPoolSize           int
	mMinIdleConns       int
	mMaxConnAge         time.Duration
	mPoolTimeout        time.Duration
	mIdleTimeout        time.Duration
	mIdleCheckFrequency time.Duration

	mRedisClient *redis.Client
	mCM iface.ConfigMap
}

func (r *Redis) Name() string {
	return Name
}

func (r *Redis) Version() string {
	return constant.Version
}

func (r *Redis) Category() string {
	return Category
}

func (r *Redis) ContractId() string {
	return ContractId
}

func (r *Redis) GetConfigMap() iface.ConfigMap {
	return r.mCM
}

func (r *Redis) New() iface.ICapability {
	return &Redis{}
}

func (r *Redis) SetConfigMap(cm iface.ConfigMap) error {
	r.mCM = cm

	// host:port address.
	r.mAddress = cm.String("redis_address", "localhost:6379")

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	r.mUserName = cm.String("redis_username", "")

	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	r.mPassword = cm.String("redis_password", "")

	// Database to be selected after connecting to the server.
	r.mDB = cm.Int("redis_db", 0)

	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	r.mUserName = cm.String("redis_username", "")

	// The network type, either tcp or unix.
	// Default is tcp.
	r.mNetwork = cm.String("redis_network", "tcp")

	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	r.mMaxRetries = cm.Int("redis_max_retries", 3)

	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	r.mMinRetryBackoff = cm.Duration("redis_min_retry_backoff", 8*time.Millisecond)

	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	r.mMaxRetryBackoff = cm.Duration("redis_max_retry_backoff", 512*time.Millisecond)

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	r.mDialTimeout = cm.Duration("redis_dialer_timeout", 5*time.Second)

	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	r.mReadTimeout = cm.Duration("redis_read_timeout", 3*time.Second)

	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	r.mWriteTimeout = cm.Duration("redis_write_timeout", 3*time.Second)

	// Type of connection pool.
	// true for FIFO pool, false for LIFO pool.
	// Note that fifo has higher overhead compared to lifo.
	r.mPoolFIFO = cm.Bool("redis_pool_fifo", true)

	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	r.mPoolSize = cm.Int("redis_pool_size", 10)

	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	r.mMinIdleConns = cm.Int("redis_min_idle_conns", 5)

	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	r.mMaxConnAge = cm.Duration("redis_max_conn_age", 0)

	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	r.mPoolTimeout = cm.Duration("redis_pool_timeout", 4*time.Second)

	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	r.mIdleTimeout = cm.Duration("redis_idle_timeout", 5*time.Minute)

	// Frequency of idle checks made by idle connections' reaper.
	// Default is 1 minute. -1 disables idle connections' reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	r.mIdleCheckFrequency = cm.Duration("redis_idle_check_frequency", 1*time.Minute)

	return nil
}

func (r *Redis) Setup() error {
	redisOpts := &redis.Options{
		Network:            r.mNetwork,
		Addr:               r.mAddress,
		Username:           r.mUserName,
		Password:           r.mPassword,
		DB:                 r.mDB,
		MaxRetries:         r.mMaxRetries,
		MinRetryBackoff:    r.mMinRetryBackoff,
		MaxRetryBackoff:    r.mMaxRetryBackoff,
		DialTimeout:        r.mDialTimeout,
		ReadTimeout:        r.mReadTimeout,
		WriteTimeout:       r.mWriteTimeout,
		PoolFIFO:           r.mPoolFIFO,
		PoolSize:           r.mPoolSize,
		MinIdleConns:       r.mMinIdleConns,
		MaxConnAge:         r.mMaxConnAge,
		PoolTimeout:        r.mPoolTimeout,
		IdleTimeout:        r.mIdleTimeout,
		IdleCheckFrequency: r.mIdleCheckFrequency,
	}

	r.mRedisClient = redis.NewClient(redisOpts)

	return nil
}

func (r *Redis) Get(ctx context.Context, key string, value interface{}) error {
	return nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&Redis{})
}
