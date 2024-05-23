package config

type Reponse[T any] struct {
	success bool
	message string
	data    []T
}
