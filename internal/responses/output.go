package responses

type Output struct {
	Message string `json:"message"`
	Body    string `json:"body"`
}

func FromError(err error) Output {
	return Output{
		Message: "Error",
		Body:    err.Error(),
	}
}
