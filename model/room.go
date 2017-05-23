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
}

func NewRoomModel(db *mgo.Database) *Room {
	room := new(Room)
	room.db = db
	room.rooms = db.C("room")
	return room
}

type RoomLog struct {
	ID         bson.ObjectId `json:"-" bson:"_id"`
	RoomID     bson.ObjectId `json:"room" bson:"room"`
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
	return room, err
}

func (r *Room) FindLog(ID bson.ObjectId, from string, till string) (roomLogs []RoomLog, err error) {
	// err = r.db.C("room_log").Find(bson.M{"room": bson.ObjectId(ID)}).One(&roomLog)
	roomLogs = []RoomLog{}
	ids := make([]bson.ObjectId, 1)
	ids[0] = ID

	/**
	* 	RoomLog
	*	{
	*	"_id" : ObjectId("592190e2932c2b231c98b511"),
	*	"room" : ObjectId("591ef303932c2b007938a88b"),
	*	"occupation" : 11,
	*	"timestamp" : NumberLong(1495371600000000000)
	*	}
	 */

	// "timestamp" : NumberLong(1495371600000000000)
	fromDate := time.Date(2017, time.May, 21, 13, 0, 0, 0, time.UTC).UnixNano()

	// "timestamp" : NumberLong(1495373400000000000)
	toDate := time.Date(2017, time.May, 21, 13, 30, 0, 0, time.UTC).UnixNano()

	log.Printf("fromDate: %s", fromDate)
	log.Printf("toDate: %s", toDate)

	err = r.db.C("room_log").Find(bson.M{"room": bson.M{"$in": ids}}).All(&roomLogs)

	// err = r.db.C("room_log").Find(bson.M{"$and": []bson.M{bson.M{"room": bson.M{"$in": ids}}, bson.M{"timestamp": bson.M{"$gte": fromDate, "$lte": toDate}}}}).All(&roomLogs)

	if err != nil {
		log.Fatal("Can not find log", err)
	}

	return roomLogs, err
}
