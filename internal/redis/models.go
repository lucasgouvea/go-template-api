package redis

type ModelMeta struct {
	DefaultScore float64
	Name         string
	Hash         string
	SortedSet    string
}

type Model[T any] struct {
	Data T
	Meta ModelMeta
}
