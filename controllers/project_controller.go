package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
	"time"
)

func ProjectIndex(db gorm.DB, r render.Render) {
	projects := []models.Project{}
	db.Find(&projects)
	r.JSON(200, projects)
}

func ProjectGet(db gorm.DB, r render.Render, params martini.Params) {
	var project models.Project
	if err := db.First(&project, params["id"]).Error; err != nil {
		r.JSON(404, map[string]interface{}{"error": "Project not found"})
		return
	}
	r.JSON(200, project)
}

func ProjectNew(r render.Render) {
	project := new(models.Project)
	r.JSON(200, project)
}

func ProjectCreate(db gorm.DB, r render.Render, project models.Project) {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	db.Save(&project)
	r.JSON(201, project)
}

func ProjectUpdate(db gorm.DB, r render.Render, params martini.Params, updatedProject models.Project) {
	var project models.Project
	if err := db.First(&project, params["id"]).Error; err != nil {
		r.JSON(404, map[string]interface{}{"error": "Project not found"})
		return
	}

	project.Name = updatedProject.Name
	project.UpdatedAt = time.Now()

	db.Save(&project)
	r.JSON(200, project)
}

func ProjectDelete(db gorm.DB, r render.Render, params martini.Params) {
	var project models.Project
	if err := db.First(&project, params["id"]).Error; err != nil {
		r.JSON(404, map[string]interface{}{"error": "Project not found"})
		return
	}
	db.Delete(project)
	r.JSON(204, nil)
}
