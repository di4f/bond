package bond

import (
	"net/http"
)

type Request = http.Request
type ResponseWriter = http.ResponseWriter
type HandlerFunc = http.HandlerFunc
type Handler = http.Handler
type Server = http.Server

type Context struct {
	R *Request
	W ResponseWriter
}

type ContextFunc func(*Context)
func (fn ContextFunc) ServeHTTP(w ResponseWriter, r *Request) {
	fn(&Context{
		R: r,
		W: w,
	})
}

