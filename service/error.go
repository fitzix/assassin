package service

import "errors"

var (
	ErrInvalidToken = errors.New("asn: invalid jwt key")
)
