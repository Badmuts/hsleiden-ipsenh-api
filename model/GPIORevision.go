package model

type GPIORevision struct {
	Name         string `json:"name"`
	RevisionCode string `json:"revisioncode"`
	Pins         int    `json:"pins"`
}
