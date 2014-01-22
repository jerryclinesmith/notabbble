package models

import (
	"labix.org/v2/mgo/bson"
)

type Project struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	//Created int64
	//Updated int64
	//OwnerId int64
}
