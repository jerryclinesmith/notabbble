package models

import (
	"labix.org/v2/mgo/bson"
	"github.com/codegangsta/martini-contrib/binding"
)

type Project struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name" form:"projectName" binding:"required"`
	//Created int64
	//Updated int64
	//OwnerId int64
}

func (p Project) Validate(errors *binding.Errors) {
	// TODO: Check for duplicate name
	if len(p.Name) < 0 {
		errors.Fields["name"] = "Name is required"
	}
}
