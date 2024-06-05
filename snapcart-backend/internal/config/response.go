package config

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

func GetResponse[T any](data []T) Response[T] {
	response := Response[T]{
		Success: true,
		Message: "Work",
		Data:    data,
	}
	return response
}