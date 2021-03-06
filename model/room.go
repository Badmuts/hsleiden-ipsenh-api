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
	HubIDs      []bson.ObjectId `json:"-" bson:"hubs,omitempty"`
	BuildingID  bson.ObjectId   `json:"building" bson:"building,omitempty"`
	db          *mgo.Database
	RoomLogs    []RoomLog       `json:"logs" bson:"-"`
	RoomRosters []RoomRoster    `json:"roster" bson:"-"`
	rooms       *mgo.Collection `bson:"-"`
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

type RoomRoster struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	RoomID       bson.ObjectId `json:"-" bson:"room"`
	PersonAmount int           `json:"amount" bson:"person_amount"`
	From         time.Time     `json:"from" bson:"from"`
	Till         time.Time     `json:"till" bson:"till"`
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

	err = r.db.C("room_log").Find(bson.M{"room": bson.ObjectId(ID)}).All(&room.RoomLogs)
	err = r.db.C("room_reservation").Find(bson.M{"room": bson.ObjectId(ID)}).All(&room.RoomRosters)

	return room, err
}

func (r *Room) Find(buildingID string) (rooms []Room, err error) {

	if bson.IsObjectIdHex(buildingID) {
		building := bson.ObjectIdHex(buildingID)
		err = r.db.C("room").Find(bson.M{"building": building}).All(&rooms)
	} else {
		err = r.rooms.Find(bson.M{}).Limit(25).All(&rooms)
	}

	for index := range rooms {
		rooms[index].db = r.db
	}

	return rooms, err
}
