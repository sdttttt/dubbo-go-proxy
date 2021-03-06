package client

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/model"
)

// Request request for endpoint
type Request struct {
	Body   []byte
	Header map[string]string
	Api    *model.Api
}

// NewRequest create a request
func NewRequest(b []byte, api *model.Api) *Request {
	return &Request{
		Body: b,
		Api:  api,
	}
}
