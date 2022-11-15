package redis

type Model[T any] struct {
	Data         T
	DefaultScore string
	Hash         string
	SortedSet    string
}
