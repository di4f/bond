package main

import (
	"github.com/omnipunk/bond"
)

var root = bond.Mux().
Def(
	"",
	bond.ContextFunc(func(c *bond.Context){
		c.W.Write([]byte("This is the index page"))
	}),
).Def(
	"hello",
	bond.Mux().Def(
		"en",
		bond.ContextFunc(func(c *bond.Context){
			c.W.Write([]byte("Hello, World!"))
		}),
	).Def(
		"ru",
		bond.ContextFunc(func(c *bond.Context){
			c.W.Write([]byte("Привет, Мир!"))
		}),
	),
)

func main() {
	srv := bond.Server{
		Addr: ":10800",
		Handler: root,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
