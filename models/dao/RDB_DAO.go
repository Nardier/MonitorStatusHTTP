package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"

	modelObject "../../models"
)

type RDB_DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "monitor"
)

func (r *RDB_DAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

func (m *RDB_DAO) GetAll() ([]modelObject.ResponseDataBody_OBJ, error) {
	var movies []modelObject.ResponseDataBody_OBJ
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *RDB_DAO) GetByID(id string) (modelObject.ResponseDataBody_OBJ, error) {
	var movie modelObject.ResponseDataBody_OBJ
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *RDB_DAO) Create(movie modelObject.ResponseDataBody_OBJ) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *RDB_DAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *RDB_DAO) Update(id string, movie modelObject.ResponseDataBody_OBJ) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)
	return err
}
