package model

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"
import "log"

type Room struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Size        float32       `json:"size" bson:"size"`
	MaxCapacity int           `json:"maxCapacity" bson:"maxCapacity"`
	Occupation  int           `json:"occupation" bson:"occupation"`
	// Hubs        []Hub           `json:"hubs" bson:"-"`
	HubIDs     []bson.ObjectId `json:"-" bson:"hubs,omitempty"`
	BuildingID bson.ObjectId   `json:"-" bson:"building"`
	db         *mgo.Database   `json:"-" bson:"-"`
	rooms      *mgo.Collection `bson:"-"`
}

func NewRoomModel(db *mgo.Database) *Room {
	room := new(Room)
	room.db = db
	room.rooms = db.C("room")
	return room
}

// Save saves room to the DB
func (r *Room) Save() (room *Room, err error) {
	if r.ID == "" {
		r.ID = bson.NewObjectId()
	}

	if _, err = r.rooms.UpsertId(r.ID, r); err != nil {
		log.Fatal("Cannot upsert room ", err)
		return nil, err
	}

	building := NewBuildingModel(r.db).FindId(r.BuildingID)
	// Todo: check if room already existst
	building.RoomIDs = append(building.RoomIDs, r.ID)

	if building, err = building.Save(); err != nil {
		log.Fatal("Cannot upsert building ", err)
		return nil, err
	}

	return r, err
}

func (r *Room) FindId(ID bson.ObjectId) (room *Room, err error) {
	err = r.rooms.FindId(ID).One(&room)
	room.db = r.db
	room.rooms = r.rooms
	return room, err
}
