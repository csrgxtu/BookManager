package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "BookManager/Server/models"
)

var BookCollection = beego.AppConfig.String("bm_mongocollection_book")

func CreateBook(book models.Book) (nBook models.Book, err error) {
  book.Id = bson.NewObjectId()
  nBook = book

  err = BMSession.DB(BMDB).C(BookCollection).Insert(&book)

  return
}

func CreateBooks(books []models.Book) (nBooks []models.Book, err error) {
  return
}

func ReadBook(book models.Book) (nBook models.Book, err error) {
  return
}

func ReadBooks(books []models.Book) (nBooks []models.Book, err error) {
  return
}

func UpdateBook(book models.Book) (nBook models.Book, err error) {
  return
}

func UpdateBooks(books []models.Book) (nBooks []models.Book, err error) {
  return
}

func DeleteBook(book models.Book) (nBook models.Book, err error) {
  return
}

func DeleteBooks(books []models.Book) (nBooks []models.Book, err error) {
  return
}
