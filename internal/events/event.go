package event

import (
	"io"
	"net/http"
)

type HTTPEndpoint struct {
	Path    string            `json:"path"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func FromRequest(req *http.Request) (HTTPEndpoint, error) {
	// Read the Body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return HTTPEndpoint{}, err
	}

	return HTTPEndpoint{
		Path:   req.URL.Path,
		Method: req.Method,
		Body:   string(bodyBytes),
	}, nil
}
