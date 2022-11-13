package redis

import (
	"context"
	"os"
	"strings"
	"time"

	Guns "go-api/internal/api/guns"

	"github.com/gomodule/redigo/redis"
)

var ctx = context.Background()

var INFINITE_TTL time.Duration = 0

func getHash(gun Guns.IGun) string {
	var hash strings.Builder
	hash.WriteString("gun:")
	hash.WriteString(gun.Id)
	return hash.String()
}

func getMany(connection redis.Conn, hashes []string) []Guns.IGun {
	var err error
	var reply interface{}
	var values []interface{}
	var guns []Guns.IGun

	for _, hash := range hashes {
		connection.Send("HGETALL", hash)
	}

	connection.Flush()

	for i := 0; i < len(Guns.Guns); i++ {
		var data Guns.IGun
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

		guns = append(guns, data)
	}

	return guns
}

func Connect() redis.Conn {
	connection, error := redis.DialContext(ctx, "tcp", os.Getenv("REDIS_ADDRESS"))

	if error != nil {
		panic(error)
	}

	return connection
}

func Setup() {

	var hashes []string
	var err error = nil
	connection := Connect()

	for _, gun := range Guns.Guns {
		var hash = getHash(gun)
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

	// TODO: move to somehwere -> var guns = getMany(connection, hashes)

	connection.Close()

}
