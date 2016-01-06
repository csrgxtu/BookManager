package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Book struct {
		Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Title        string        `json:"title" bson:"title"`
    Isbn         string        `json:"isbn13" bson:"isbn13"`
		Image	       string        `json:"image" bson:"image"`
    Author       []string      `json:"author" bson:"author"`
    Publisher    string        `json:"publisher" bson:"publisher"`
	}
)
