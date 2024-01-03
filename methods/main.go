package methods

import (
	"net/http"
)

type Method string

const (
	Get Method = http.MethodGet
	Post Method = http.MethodPost
	Put Method = http.MethodPut
	Head Method = http.MethodHead
	Patch Method = http.MethodPatch
	Delete Method = http.MethodDelete
	Connect Method = http.MethodConnect
	Options Method = http.MethodOptions
	Trace Method = http.MethodTrace
)

func (m Method) String() string {
	return string(m)
}

