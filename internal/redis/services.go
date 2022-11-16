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
	defer connection.Close()

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

			models = append(models, Model[T]{Data: data})
		}
	}

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

func CreateOne[T any](model *Model[T]) {
	var connection = GetConnection()
	defer connection.Close()

	var argsHMSET = redis.Args{}.Add(model.Meta.Hash).AddFlat(&model.Data)
	var argsZADD = redis.Args{}.Add(model.Meta.SortedSet).Add(model.Meta.DefaultScore).Add(model.Meta.Hash)
	var err error
	var reply any

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

	if _, err = connection.Receive(); err != nil {
		panic(err)
	}

}
