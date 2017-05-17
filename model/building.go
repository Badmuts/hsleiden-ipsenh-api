package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Building struct {
	ID        bson.ObjectId   `json:"id" bson:"_id"`
	Name      string          `json:"name" bson:"name"`
	RoomIDs   []bson.ObjectId `json:"-" bson:"rooms"`
	Location  string          `json:"location"`
	db        *mgo.Database
	buildings *mgo.Collection
}

func NewBuildingModel(db *mgo.Database) *Building {
	building := new(Building)
	building.db = db
	building.buildings = db.C("building")
	return building
}

func (b *Building) Save() (building *Building, err error) {
	if b.ID == "" {
		b.ID = bson.NewObjectId()
		err = b.buildings.Insert(b)
	} else {
		err = b.buildings.UpdateId(b.ID, b)
	}
	return b, err
}

func (b *Building) FindId(ID bson.ObjectId) *Building {
	building := NewBuildingModel(b.db)
	b.buildings.FindId(ID).One(&building)
	building.db = b.db
	building.buildings = b.db.C("building")
	return building
}

func (b *Building) GetRooms() (rooms []Room, err error) {
	rooms = make([]Room, 0)
	err = b.db.C("room").Find(bson.M{"_id": bson.M{"$in": b.RoomIDs}}).All(&rooms)
	return rooms, err
}
