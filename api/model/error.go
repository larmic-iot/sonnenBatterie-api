package model

type ProtocolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
