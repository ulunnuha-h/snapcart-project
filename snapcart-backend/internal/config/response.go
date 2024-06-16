package config

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

func GetResponse[T any](data []T, status bool, message string) Response[T] {
	response := Response[T]{
		Success: status,
		Message: message,
		Data:    data,
	}
	return response
}