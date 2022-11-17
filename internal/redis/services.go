package redis

func List[T any](modelName string) *[]Model[T] {
	const offset = 0
	const limit = 10

	var connection = GetConnection()
	defer connection.Close()

	hashes := zRange(connection, modelName, offset, limit)
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
