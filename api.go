package bond

import (
	//"io"
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
)

type ContentType string

const (
	PlainText ContentType = "text/plain; charset=utf-8"
)

type Decoder interface {
	Decode(any) error
}

type ApiFunc func(*Context)
func (fn ApiFunc) ServeHTTP(w ResponseWriter, r *Request) {
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
	dec Decoder
}

func (c *Context) SetContentType(typ ContentType) {
	c.SetHeader("Content-Type", string(typ))
}

func (c *Context) ContentType() string {
	ret, ok := c.Header("Content-Type")
	if !ok {
		return ""
	}
	if len(ret) < 1 {
		return ""
	}
	return ret[0]
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
func (c *Context) Scan(v any) error {
	if c.dec == nil {
		typ := c.ContentType()
		switch typ {
		case "application/json" :
			c.dec = json.NewDecoder(c.R.Body)
		default:
			return UnknownContentTypeErr
		}
	}
	err := c.dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
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
