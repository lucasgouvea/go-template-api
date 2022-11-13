package redis

import (
	"flag"

	Shared "go-api/internal/shared"

	"github.com/gomodule/redigo/redis"
)

func GetMany[T any](hashes []string) []T {
	var err error
	var reply interface{}
	var values []interface{}
	var models []T
	var connection = GetConnection()

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

		if len(values) > 0 {
			err = redis.ScanStruct(values, &data)
			if err != nil {
				panic(err)
			}

			models = append(models, data)
		}
	}

	connection.Close()

	return models
}

func CreateMany[T any](models []Shared.Model[T]) {
	var err error = nil

	flag.Parse()

	var connection = GetConnection()

	for _, model := range models {
		var args = redis.Args{}.Add(model.Hash).AddFlat(&model.Data)
		connection.Send("HMSET", args...)
	}

	connection.Flush()

	for i := 0; i < len(models); i++ {
		_, err = connection.Receive()
		if err != nil {
			panic(err)
		}
	}

	connection.Close()
}
