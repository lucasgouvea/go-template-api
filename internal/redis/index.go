package redis

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var INFINITE_TTL time.Duration = 0
var pool *redis.Pool

type IBaseModel[T any] interface {
	GetData() []T
	GetHash() string
}
type Model[T any] struct {
	Data T
	Hash string
}

func NewPool() {

	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", os.Getenv("REDIS_ADDRESS")) },
	}
}

func GetConnection() redis.Conn {
	connection := pool.Get()
	return connection
}
