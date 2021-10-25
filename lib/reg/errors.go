package reg

import (
	"errors"
)

var (
	errFound       = errors.New("found")
	errNilHandler  = errors.New("nil handler")
	errNilReceiver = errors.New("nil receiver")
	errNotFound    = errors.New("not found")
)
