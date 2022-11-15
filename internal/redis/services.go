package redis

import (
	"github.com/gomodule/redigo/redis"
)

/* func GetMany[T any]() []Model[T] {
	var connection = GetConnection()
} */

func GetManyByHashes[T any](hashes []string) []Model[T] {
	var err error
	var reply any
	var values []any
	var models []Model[T]
	var connection = GetConnection()

	for _, hash := range hashes {
		connection.Send("HGETALL", hash)
	}

	connection.Flush()

	for i := 0; i < len(hashes); i++ {
		var data T
		if reply, err = connection.Receive(); err != nil {
			panic(err)
		}

		if values, err = redis.Values(reply, err); err != nil {
			panic(err)
		}

		if len(values) > 0 {
			if err = redis.ScanStruct(values, &data); err != nil {
				panic(err)
			}

			models = append(models, Model[T]{Data: data, Hash: hashes[i]})
		}
	}

	connection.Close()

	return models
}

func GetOne[T any](hash string) *Model[T] {
	var data T
	var values []interface{}
	var connection = GetConnection()
	var reply, err = connection.Do("HGETALL", hash)
	connection.Close()
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

		var model = Model[T]{Data: data, Hash: hash}
		return &model
	}

	return nil
}

func CreateMany[T any](models []Model[T]) {
	var err error = nil

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

func CreateOne[T any](model *Model[T]) *Model[T] {
	var connection = GetConnection()
	var argsHMSET = redis.Args{}.Add(model.Hash).AddFlat(&model.Data)
	var argsZADD = redis.Args{}.Add(model.SortedSet).Add(model.DefaultScore).Add(model.Hash)
	var err error
	var reply any
	var values []any
	var data T

	if err = connection.Send("HMSET", argsHMSET...); err != nil {
		panic(err)
	}

	if err = connection.Send("ZADD", argsZADD...); err != nil {
		panic(err)
	}

	connection.Flush()

	if reply, err = connection.Receive(); err != nil || reply != "OK" {
		panic(err)
	}

	if reply, err = connection.Receive(); err != nil || reply != "OK" {
		panic(err)
	}

	if values, err = redis.Values(reply, err); err != nil {
		panic(err)
	}

	if len(values) > 0 {
		err = redis.ScanStruct(values, &data)
		if err != nil {
			panic(err)
		}

		var model = Model[T]{Data: data, Hash: model.Hash}
		return &model
	}

	connection.Close()

	return nil
}
