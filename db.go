package main

import (
	"github.com/codegangsta/martini"
	"labix.org/v2/mgo"
)

func DB() martini.Handler {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB("notabbble-dev"))

		defer s.Close()
		c.Next()
	}
}
