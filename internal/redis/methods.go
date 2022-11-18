package redis

import "github.com/gomodule/redigo/redis"

func hgetAll[T any](connection redis.Conn, hashes []string) []Model[T] {
	var err error
	var reply any
	var values []any
	var data T
	var models []Model[T]

	for _, hash := range hashes {
		connection.Send("HGETALL", hash)
	}

	connection.Flush()

	for i := 0; i < len(hashes); i++ {
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

func zRange(connection redis.Conn, modelName string, offset int, limit int) []string {
	var reply interface{}
	var hashes []string
	var err error
	var argsZRANGE = redis.Args{}.Add(modelName).Add(offset).Add(limit)

	if reply, err = connection.Do("ZRANGE", argsZRANGE...); err != nil {
		panic(err)
	}

	if hashes, err = redis.Strings(reply, err); err != nil {
		panic(err)
	}

	return hashes
}

func create[T any](connection redis.Conn, model *Model[T]) bool {
	var err error
	var argsHMSET = redis.Args{}.Add(model.Meta.Hash).AddFlat(&model.Data)
	var argsZADD = redis.Args{}.Add(model.Meta.SortedSet).Add("NX").Add(model.Meta.DefaultScore).Add(model.Meta.Hash)
	var replies []any
	var alreadyExistsCode int64 = 0
	var commands = []RedisCommand{
		{name: "MULTI"},
		{name: "ZADD", args: argsZADD},
		{name: "HMSET", args: argsHMSET},
		{name: "EXEC"},
	}

	if err = sendMany(connection, commands); err != nil {
		panic(err)
	}

	connection.Flush()

	if replies, err = receiveAll(connection, len(commands)); err != nil {
		panic(err)
	}

	var reply []any = replies[3].([]any)

	if reply[0] != alreadyExistsCode && reply[1] == "OK" {
		return true
	}

	return false
}
