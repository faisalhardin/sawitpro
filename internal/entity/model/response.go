package model

type Response struct {
	Data  string `json:"message"`
	Error string `json:"error,omitempty"`
}
