package main

import (
	"github.com/di4f/bond"
	"github.com/di4f/bond/methods"
	"github.com/di4f/bond/contents"
)

type GetNotesOptions struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var root = bond.Mux().
	Def(
		"",
		bond.Func(func(c *bond.Context) {
			c.W.Write([]byte("This is the index page"))
		}),
	).Def(
	"hello",
	bond.Mux().Def(
		"en",
		bond.Func(func(c *bond.Context) {
			c.Printf("Hello, World!")
		}),
	).Def(
		"ru",
		bond.Func(func(c *bond.Context) {
			c.Printf("Привет, Мир!")
		}),
	),
).Def(
	"web",
	bond.Static("./static"),
).Def(
	"test",
	bond.Func(func(c *bond.Context) {
		c.SetContentType(contents.Plain)
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
).Def(
	"get-notes",
	bond.Method().Def(
		methods.Get,
		bond.Func(func(c *bond.Context){
			opts := GetNotesOptions{}
			c.Scan(&opts)
			c.Printf("%v", opts)
		}),
	),
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
