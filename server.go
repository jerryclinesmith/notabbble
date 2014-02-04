package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jerryclinesmith/notabbble/controllers"
	"github.com/jerryclinesmith/notabbble/db"
	"github.com/jerryclinesmith/notabbble/models"
)

func main() {
	m := martini.Classic()
	m.Map(db.Connect())

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "home", "")
	})
	m.Get("/projects", controllers.ProjectIndex)
	m.Get("/projects/new", controllers.ProjectNew)
	m.Post("/projects", binding.Bind(models.Project{}), controllers.ProjectCreate)

	m.Run()
}
