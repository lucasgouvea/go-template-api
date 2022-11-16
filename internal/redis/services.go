package redis

import (
	"github.com/gomodule/redigo/redis"
)

func List[T any](modelName string) any {
	const offset = 0
	const limit = 10

	var connection = GetConnection()

	defer connection.Close()

	hashes := zRange(connection, modelName, offset, limit)
	models := hgetAll[T](connection, hashes)

	return models
}

func GetManyByHashes[T any](hashes []string) []Model[T] {
	var connection = GetConnection()
	defer connection.Close()

	models := hgetAll[T](connection, hashes)

	return models
}

func GetOne[T any](hash string) *Model[T] {
	var model Model[T]
	var values []interface{}
	var connection = GetConnection()
	defer connection.Close()
	var reply, err = connection.Do("HGETALL", hash)

	if err != nil {
		panic(err)
	}

	values, err = redis.Values(reply, err)
	if err != nil {
		panic(err)
	}

	if len(values) > 0 {
		err = redis.ScanStruct(values, &model.Data)
		if err != nil {
			panic(err)
		}

		var model = Model[T]{Data: model.Data}
		return &model
	}

	return nil
}

func CreateMany[T any](models []Model[T]) {
	var err error = nil

	var connection = GetConnection()
	defer connection.Close()

	for _, model := range models {
		var args = redis.Args{}.Add(model.Meta.Hash).AddFlat(&model.Data)
		connection.Send("HMSET", args...)
	}

	connection.Flush()

	for i := 0; i < len(models); i++ {
		_, err = connection.Receive()
		if err != nil {
			panic(err)
		}
	}
}

func CreateOne[T any](model *Model[T]) bool {
	var connection = GetConnection()
	defer connection.Close()

	var argsHMSET = redis.Args{}.Add(model.Meta.Hash).AddFlat(&model.Data)
	var argsZADD = redis.Args{}.Add(model.Meta.SortedSet).Add("NX").Add(model.Meta.DefaultScore).Add(model.Meta.Hash)
	var err error
	var reply any
	var alreadyExistsCode int64 = 0

	if err = connection.Send("ZADD", argsZADD...); err != nil {
		panic(err)
	}

	if err = connection.Send("HMSET", argsHMSET...); err != nil {
		panic(err)
	}

	connection.Flush()

	if reply, err = connection.Receive(); err != nil {
		panic(err)
	}

	if reply == alreadyExistsCode {
		return false
	}

	if reply, err = connection.Receive(); err != nil || reply != "OK" {
		panic(err)
	}

	return true
}
