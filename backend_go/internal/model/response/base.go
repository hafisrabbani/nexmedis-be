package response

type BaseResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{
		Status:  true,
		Message: "success",
		Data:    data,
	}
}

func Error(message string) BaseResponse[any] {
	return BaseResponse[any]{
		Status:  false,
		Message: message,
	}
}

func Ok() BaseResponse[any] {
	return BaseResponse[any]{
		Status:  true,
		Message: "success",
	}
}
