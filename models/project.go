package models

import (
	"github.com/codegangsta/martini-contrib/binding"
	"time"
)

type Project struct {
	Id        int64
	Name      string `sql:"size:255;not null;unique" form:"projectName" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Project) Validate(errors *binding.Errors) {
	// TODO: Check for duplicate name
	if len(p.Name) < 0 {
		errors.Fields["name"] = "Name is required"
	}
}
