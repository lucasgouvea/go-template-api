package redis

import (
	Shared "go-api/internal/shared"
)

func List[T any](modelName string, query Shared.Query) *[]Model[T] {

	var connection = GetConnection()
	defer connection.Close()

	hashes := zRange(connection, modelName, query.Offset, query.Limit-1)
	models := hgetAll[T](connection, hashes)

	return &models
}

func GetManyByHashes[T any](hashes []string) []Model[T] {
	var connection = GetConnection()
	defer connection.Close()

	models := hgetAll[T](connection, hashes)

	return models
}

func GetOne[T any](hash string) *Model[T] {
	var connection = GetConnection()
	defer connection.Close()
	models := hgetAll[T](connection, []string{hash})

	if len(models) == 0 {
		return nil
	}

	return &models[0]
}

func CreateOne[T any](model *Model[T]) bool {
	var connection = GetConnection()
	defer connection.Close()

	created := create(connection, model)

	return created
}
