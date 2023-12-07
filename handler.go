package bond

import (
	"net/http"
)

type Request = http.Request
type ResponseWriter = http.ResponseWriter
type HandlerFunc = http.HandlerFunc
type Handler = http.Handler
type Server = http.Server

func Static(pth string) Handler {
	return http.FileServer(http.Dir(pth))
}
