package model

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
	RoomLogs   []RoomLog `json:"logs" bson:"-"`
}

func NewRoomModel(db *mgo.Database) *Room {
	room := new(Room)
	room.db = db
	room.rooms = db.C("room")
	return room
}

type RoomLog struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	RoomID     bson.ObjectId `json:"-" bson:"room"`
	Occupation int           `json:"occupation" bson:"occupation"`
	Timestamp  time.Time     `json:"time" bson:"timestamp"`
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

	ids := make([]bson.ObjectId, 1)
	ids[0] = ID

	err = r.db.C("room_log").Find(bson.M{"room": bson.M{"$in": ids}}).All(&room.RoomLogs)

	return room, err
}
