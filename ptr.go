package goptr

func Ptr[T any](v T) *T {
	return &v
}

func P[T any](v T) *T {
	return Ptr(v)
}
