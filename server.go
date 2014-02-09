package notabbble

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jerryclinesmith/notabbble/controllers"
	"github.com/jerryclinesmith/notabbble/db"
	"github.com/jerryclinesmith/notabbble/models"
)

func InitServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Map(db.Connect())

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "home", "")
	})
	m.Get("/api/projects", controllers.ProjectIndex)
	m.Get("/api/project/:id", controllers.ProjectGet)
	m.Get("/api/projects/new", controllers.ProjectNew)
	m.Post("/api/projects", binding.Bind(models.Project{}), controllers.ProjectCreate)
	m.Put("/api/projects/:id", binding.Bind(models.Project{}), controllers.ProjectCreate)
	m.Delete("/api/projects/:id", controllers.ProjectCreate)

	return m
}

func main() {
	m := InitServer()
	m.Run()
}
