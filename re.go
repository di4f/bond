package bond

import (
	//"regexp"
)

/*type ReChecker struct {
	re *regexp.Regexp
	handler ContextHandler
}

func Re(re string, handler ContextHandler) *ReChecker {
	ret := &ReChecker{}
	ret.re = regexp.MustCompile(re)
	ret.handler = handler
	return ret
}

func (rr *ReChecker) ServeHttp(c *Context) {
	match := rr.re.MatchString(c.Path())
	if !match {
		c.NotFound()
		return
	}
	c.re = rr.re
	c.reSubmatches = c.re.FindStringSubmatch(c.Path())
	rr.handler.ServeHttp(c)
}*/

