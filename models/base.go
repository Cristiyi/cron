package models

import (
	"cron/conf"
	"log"

	"gopkg.in/mgo.v2"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

var sess *mgo.Session
var db *mgo.Database

func init() {
	SetupDB()
}

func SetupDB() *mgo.Database {
	config := conf.ReadConfig()
	sess, err := mgo.Dial(config.DB.URL)
	CheckErr(err)
	sess.SetSafe(&mgo.Safe{})
	db = sess.DB("cron")
	return db
}

func CloseDB() {
	sess.Close()
}
