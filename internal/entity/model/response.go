package model

import "fmt"

type Response struct {
	Code    int    `json:"-"`
	ErrName string `json:"error,omitempty"`
	Data    string `json:"message,omitempty"`
}

func (resp *Response) Error() string { return fmt.Sprintf("%s: %s", resp.ErrName, resp.Data) }
