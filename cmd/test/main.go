package main

import (
	"github.com/di4f/bond"
)

var root = bond.Mux().
	Def(
		"",
		bond.ApiFunc(func(c *bond.Context) {
			c.W.Write([]byte("This is the index page"))
		}),
	).Def(
	"hello",
	bond.Mux().Def(
		"en",
		bond.ApiFunc(func(c *bond.Context) {
			c.W.Write([]byte("Hello, World!"))
		}),
	).Def(
		"ru",
		bond.ApiFunc(func(c *bond.Context) {
			c.W.Write([]byte("Привет, Мир!"))
		}),
	),
).Def(
	"web",
	bond.Static("./static"),
).Def(
	"test",
	bond.ApiFunc(func(c *bond.Context) {
		c.SetContentType(bond.PlainText)
		c.Printf(
			"Path: %q\n"+
				"Content-Type: %q\n",
			c.Path(), c.ContentType(),
		)
		c.Printf("Query:\n")
		for k, vs := range c.Query() {
			c.Printf("\t%q:\n", k)
			for _, v := range vs {
				c.Printf("\t\t%q\n", v)
			}
		}
	}),
)

func main() {
	srv := bond.Server{
		Addr:    ":10800",
		Handler: root,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
