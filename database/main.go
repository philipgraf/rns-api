package database

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/philipgraf/rns-api/config"
	"gopkg.in/mgo.v2"
)

type key int

const contextKey key = 0

func Connect() (*mgo.Database, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, err
	}
	session, err2 := mgo.Dial(conf.DB.URL)
	if err2 != nil {
		return nil, err2
	}
	//defer session.Close()

	return session.DB(conf.DB.Name), nil
}

func GetFromContext(r *http.Request) *mgo.Database {
	if rv := context.Get(r, contextKey); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}

func SetToContext(r *http.Request, db *mgo.Database) {
	context.Set(r, contextKey, db)
}
