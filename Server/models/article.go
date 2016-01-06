package models

import (
	// bson "labix.org/v2/mgo/bson"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Article struct {
	Id                  bson.ObjectId `json:"id" bson:"_id"`
	Title               string        `json:"title" bson:"title"`
	Author              string        `json:"author" bson:"author"`
	Tags                []string      `json:"tags" bson:"tags"`
	Content             string        `json:"content" bson:"content"`
	Summary							string				`json:"summary" bson:"summary"`
	Creator             string        `json:"creator" bson:"creator"`
	Publisher           string        `json:"publisher" bson:"publisher"`
	PublicationPosition int           `json:"publication_position" bson:"publication_position"`
	PublicationStatus   string        `json:"publication_status" bson:"publication_status"`
	DateCreated         time.Time     `json:"date_created" bson:"date_created"`
	DateModified        time.Time     `json:"date_modified" bson:"date_modified"`
	DateStartPublish    time.Time     `json:"date_start_publish" bson:"date_start_publish"`
	DateEndPublish      time.Time     `json:"date_end_publish" bson:"date_end_publish"`
	FavImg							string				`json:"fav_img" bson:"fav_img"`
	RecArticles					[]string			`json:"rec_articles" bson:"rec_articles"`
	RelatedAuthors			[]string			`json:"related_authors" bson:"related_authors"`
	ArticleRank					string				`json:"article_rank" bson:"article_rank"`
	Status							bool					`json:"status" bson:"status"`
	Number							int64					`json:"number" bson:"number"`
	QiniuUrl						string				`json:"qiniuurl" bson:"qiniuurl"`
}
