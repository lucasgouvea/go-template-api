package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	Guns "go-api/internal/api/guns"

	"github.com/gomodule/redigo/redis"
)

type Rediser[T any] interface {
	getConnection() redis.Conn
	GetMany(hashes []string) []T
}

var ctx = context.Background()
var INFINITE_TTL time.Duration = 0
var Connection redis.Conn = nil

func GetMany[T any](hashes []string) []T {
	var err error
	var reply interface{}
	var values []interface{}
	var models []T
	var connection = getConnection()

	for _, hash := range hashes {
		connection.Send("HGETALL", hash)
	}

	connection.Flush()

	for i := 0; i < len(hashes); i++ {
		var data T
		reply, err = connection.Receive()
		if err != nil {
			panic(err)
		}
		values, err = redis.Values(reply, err)
		if err != nil {
			panic(err)
		}
		err = redis.ScanStruct(values, &data)
		if err != nil {
			panic(err)
		}

		models = append(models, data)
	}

	return models
}

func getConnection() redis.Conn {
	if Connection == nil {
		connection, error := redis.DialContext(ctx, "tcp", os.Getenv("REDIS_ADDRESS"))
		if error != nil {
			panic(error)
		}
		Connection = connection
	}
	return Connection
}

func Setup() {
	var hashes []string
	var err error = nil
	connection := getConnection()

	for _, gun := range Guns.Guns {
		var hash = gun.GetHash()
		hashes = append(hashes, hash)
		connection.Send("HMSET", redis.Args{}.Add(hash).AddFlat(&gun)...)
	}

	connection.Flush()

	for i := 0; i < len(Guns.Guns); i++ {
		_, err = connection.Receive()
		if err != nil {
			panic(err)
		}
	}

	var guns = GetMany[Guns.IGun](hashes)
	fmt.Printf("Initialized guns: \n")
	for _, gun := range guns {
		fmt.Printf("%v\n", gun.Name)
	}

	connection.Close()
}
