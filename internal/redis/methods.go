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

func zRange(connection redis.Conn, modelName string, offset int64, limit int64) []string {
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
