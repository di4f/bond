package bond

import (
	"net/http"
	"strings"
	//"fmt"
	"path"
)

type Router struct {
	pathMap map[string] Handler
}

func Mux() *Router {
	ret := &Router{}
	ret.pathMap = map[string] Handler{}
	return ret
}

func (router *Router) Def(pth string, handler Handler) *Router {
	_, dup := router.pathMap[pth]
	if dup {
		panic(DupDefErr)
	}
	router.pathMap[pth] = handler
	return router
}

func (router *Router) ServeHTTP(w ResponseWriter, r *Request) {
	pth := r.URL.Path
	pth = path.Clean(pth)
	pths := strings.SplitN(pth, "/", 3)

	var name string
	if len(pths) > 1 {
		name = pths[1]
	}
	name, _ = strings.CutSuffix(name, "/")

	prefix := "/"
	if pth != "/" {
		prefix = path.Clean("/" + name)
	}

	//fmt.Printf("Path: %q\n", r.URL.Path)
	//fmt.Printf("%q %q %q %q\n", pth, prefix, pths, name)
	handler, ok := router.pathMap[name]
	if !ok {
		http.NotFound(w, r)
		return
	}

	r.URL.Path = pth
	http.StripPrefix(prefix, handler).ServeHTTP(w, r)
}

