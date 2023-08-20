package dbredis

import "time"

type Channels map[string]Config

type Config struct {
	Addrs                 []string      `mapstructure:"addrs" json:"addrs,omitempty" bson:"addrs,omitempty"`
	ClientName            string        `mapstructure:"client_name" json:"client_name,omitempty" bson:"client_name,omitempty"`
	DB                    int           `mapstructure:"db" json:"db,omitempty" bson:"db,omitempty"`
	Username              string        `mapstructure:"username" json:"username,omitempty" bson:"username,omitempty"`
	Password              string        `mapstructure:"password" json:"password,omitempty" bson:"password,omitempty"`
	SentinelUsername      string        `mapstructure:"sentinel_username" json:"sentinel_username,omitempty" bson:"sentinel_username,omitempty"`
	SentinelPassword      string        `mapstructure:"sentinel_password" json:"sentinel_password,omitempty" bson:"sentinel_password,omitempty"`
	MaxRetries            int           `mapstructure:"max_retries" json:"max_retries,omitempty" bson:"max_retries,omitempty"`
	MinRetryBackoff       time.Duration `mapstructure:"min_retry_backoff" json:"min_retry_backoff,omitempty" bson:"min_retry_backoff,omitempty"`
	MaxRetryBackoff       time.Duration `mapstructure:"max_retry_backoff" json:"max_retry_backoff,omitempty" bson:"max_retry_backoff,omitempty"`
	DialTimeout           time.Duration `mapstructure:"dial_timeout" json:"dial_timeout,omitempty" bson:"dial_timeout,omitempty"`
	ReadTimeout           time.Duration `mapstructure:"read_timeout" json:"read_timeout,omitempty" bson:"read_timeout,omitempty"`
	WriteTimeout          time.Duration `mapstructure:"write_timeout" json:"write_timeout,omitempty" bson:"write_timeout,omitempty"`
	ContextTimeoutEnabled bool          `mapstructure:"context_timeout_enabled" json:"context_timeout_enabled,omitempty" bson:"context_timeout_enabled,omitempty"`
	PoolFIFO              bool          `mapstructure:"pool_fifo" json:"pool_fifo,omitempty" bson:"pool_fifo,omitempty"`
	PoolSize              int           `mapstructure:"pool_size" json:"pool_size,omitempty" bson:"pool_size,omitempty"`
	PoolTimeout           time.Duration `mapstructure:"pool_timeout" json:"pool_timeout,omitempty" bson:"pool_timeout,omitempty"`
	MinIdleConns          int           `mapstructure:"min_idle_conns" json:"min_idle_conns,omitempty" bson:"min_idle_conns,omitempty"`
	MaxIdleConns          int           `mapstructure:"max_idle_conns" json:"max_idle_conns,omitempty" bson:"max_idle_conns,omitempty"`
	ConnMaxIdleTime       time.Duration `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time,omitempty" bson:"conn_max_idle_time,omitempty"`
	ConnMaxLifetime       time.Duration `mapstructure:"conn_max_life_time" json:"conn_max_lifetime,omitempty" bson:"conn_max_lifetime,omitempty"`
	MaxRedirects          int           `mapstructure:"max_redirects" json:"max_redirects,omitempty" bson:"max_redirects,omitempty"`
	ReadOnly              bool          `mapstructure:"read_only" json:"read_only,omitempty" bson:"read_only,omitempty"`
	RouteByLatency        bool          `mapstructure:"route_by_latency" json:"route_by_latency,omitempty" bson:"route_by_latency,omitempty"`
	RouteRandomly         bool          `mapstructure:"route_randomly" json:"route_randomly,omitempty" bson:"route_randomly,omitempty"`
	MasterName            string        `mapstructure:"master_name" json:"master_name,omitempty" bson:"master_name,omitempty"`
	Ping                  bool          `mapstructure:"ping" json:"ping,omitempty" bson:"ping,omitempty"`
}
