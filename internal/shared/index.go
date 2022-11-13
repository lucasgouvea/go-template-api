package shared

type Model[T any] struct {
	Data T
	Hash string
}
