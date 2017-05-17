package model

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

type Room struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Size        float32       `json:"size" bson:"size"`
	MaxCapacity int           `json:"maxCapacity" bson:"maxCapacity"`
	Occupation  int           `json:"occupation" bson:"occupation"`
	// Hubs        []Hub           `json:"hubs" bson:"-"`
	HubIDs     []bson.ObjectId `json:"-" bson:"hubs"`
	BuildingID bson.ObjectId   `json:"-" bson:"building"`
	db         *mgo.Database
	rooms      *mgo.Collection
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

	_, err = r.rooms.Upsert(r, r)

	// Todo: save building relation
	building := NewBuildingModel(r.db).FindId(r.BuildingID)
	if building.RoomIDs == nil {
		building.RoomIDs = []bson.ObjectId{}
	}
	building.RoomIDs = append(building.RoomIDs, r.ID)
	building.Save()

	return r, err
}
