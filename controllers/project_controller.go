package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
)

func ProjectIndex(db gorm.DB, r render.Render) {
	var projects []models.Project
	db.Find(&projects)
	r.HTML(200, "projects/index", projects)
}

func ProjectNew(r render.Render) {
	project := new(models.Project)
	r.HTML(200, "projects/new", project)
}

func ProjectCreate(db gorm.DB, r render.Render, p models.Project) {
	db.Save(p)
	r.Redirect("/projects")
}
