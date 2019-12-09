package library

import "github.com/go-redis/redis/v7"

var (
	RedisConn *redis.Conn

	HOST = "127.0.0.1:"
	PORT = "6379"
	AUTH = ""
)

func InitRedis() {
	Addr := HOST + PORT
	RedisConn = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               Addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           AUTH,
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}).Conn()
}
