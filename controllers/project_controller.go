package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jerryclinesmith/notabbble/models"
	"labix.org/v2/mgo"
)

func ProjectIndex(db *mgo.Database, r render.Render) {
	var projects []models.Project
	err := db.C("projects").Find(nil).All(&projects)
	if err != nil {
		panic(err)
	}
	r.HTML(200, "projects/index", projects)
}

func ProjectNew(r render.Render) {
	project := new(models.Project)
	r.HTML(200, "projects/new", project)
}

func ProjectCreate(db *mgo.Database, r render.Render, p models.Project) {
	err := db.C("projects").Insert(p)
	if err != nil {
		panic(err)
	}
	r.Redirect("/projects")
}
