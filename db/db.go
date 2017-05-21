package db

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Connect returns connection to Database
func Connect() *mgo.Database {
	host := os.Getenv("MONGO_HOST")
	db := os.Getenv("MONGO_DB")

	if host == "" {
		host = "db"
	}
	if db == "" {
		db = "ipsenh"
	}

	// mgo.SetDebug(true)

	// var aLogger *log.Logger
	// aLogger = log.New(os.Stderr, "", log.LstdFlags)
	// mgo.SetLogger(aLogger)

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	// defer session.Close()
	return session.DB(db)
}
