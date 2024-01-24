package model

import (
	"context"
)

// AppContext ...
type AppContext struct {
	context.Context
	MandatoryRequest
	Token  string `json:"token"`
	UID    string `json:"uid"`
	Issuer string `json:"iss"`
}

// MandatoryRequest ...
type MandatoryRequest struct {
	ChannelID string
}
