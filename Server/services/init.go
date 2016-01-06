package services

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var BMSession *mgo.Session
var BMDB = beego.AppConfig.String("bm_mongodb")

func init() {
  BMSessionInit()
}

func BMSessionInit() {
  var err error

  BMSession, err = mgo.Dial(beego.AppConfig.String("bm_mongourl"))
  if err != nil {
    panic(err)
  }

  BMSession.SetMode(mgo.Monotonic, true)
}
