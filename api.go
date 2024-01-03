package bond

import (
	"io"
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
	"github.com/di4f/bond/contents"
)

type Decoder interface {
	Decode(any) error
}

type Func func(*Context)
func (fn Func) ServeHTTP(w ResponseWriter, r *Request) {
	fn(&Context{
		R: r,
		W: w,
	})
}

type Context struct {
	R *Request
	W ResponseWriter
	// Custom data to store stuff.
	Data any

	scanErr error
	dec Decoder
}

func (c *Context) SetContentType(typ contents.Type) {
	c.SetHeader("Content-Type", string(typ))
}

func (c *Context) ContentType() contents.Type {
	ret, ok := c.Header("Content-Type")
	if !ok {
		return ""
	}
	if len(ret) < 1 {
		return ""
	}
	return contents.Type(ret[0])
}

func (c *Context) SetHeader(k, v string) {
	c.W.Header().Set(k, v)
}

func (c *Context) Header(name string) ([]string, bool) {
	ret, ok := c.R.Header[name]
	return ret, ok
}

// Closes the requests body after finishes scaning.
func (c *Context) Close() {
	c.R.Body.Close()
}

// Scan the incoming value from body depending
// on the content type of the request.
func (c *Context) Scan(v any) bool {
	if c.dec == nil {
		typ := c.ContentType()
		switch typ {
		case contents.Json :
			c.dec = json.NewDecoder(c.R.Body)
		default:
			c.scanErr = UnknownContentTypeErr
			return false
		}
	}
	err := c.dec.Decode(v)
	if err != nil {
		if err != io.EOF {
			c.scanErr = err
		}
		return false
	}
	return true
}

func (c *Context) ScanErr() error {
	return c.scanErr
}

func (c *Context) Path() string {
	return c.R.URL.Path
}

func (c *Context) NotFound() {
	http.NotFound(c.W, c.R)
}

func (c *Context) Printf(format string, v ...any) (int, error) {
	return fmt.Fprintf(c.W, format, v...)
}

func (c *Context) Query() url.Values {
	return c.R.URL.Query()
}
