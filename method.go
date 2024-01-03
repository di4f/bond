package bond

import (
	"net/http"
	"github.com/di4f/bond/methods"
)

// The type implements functionality for multiplexing
// the methods.
type MethodRouter struct {
	methodMap map[methods.Method] Handler
}

func Method() *MethodRouter {
	ret := &MethodRouter{}
	ret.methodMap = make(map[methods.Method] Handler)
	return ret
}

func (mr *MethodRouter) Def(
	method methods.Method,
	handler Handler,
) *MethodRouter {
	_, dup := mr.methodMap[method]
	if dup {
		panic("got the duplicate method")
	}

	mr.methodMap[method] = handler
	return mr
}

func (mr *MethodRouter) ServeHTTP(w ResponseWriter, r *Request) {
	handler, ok := mr.methodMap[methods.Method(r.Method)]
	if !ok {
		http.NotFound(w, r)
		return
	}
	handler.ServeHTTP(w, r)
}
