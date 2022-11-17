package redis

import "github.com/gomodule/redigo/redis"

type RedisCommand struct {
	name string
	args redis.Args
}

func sendMany(connection redis.Conn, commands []RedisCommand) error {
	var err error
	for _, command := range commands {
		if len(command.args) > 0 {
			err = connection.Send(command.name, command.args...)
		} else {
			err = connection.Send(command.name)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func receiveAll(connection redis.Conn, count int) ([]any, error) {
	var reply any
	var err error
	var replies []any
	for i := 0; i < count; i++ {
		reply, err = connection.Receive()
		if err != nil {
			return []any{}, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}
