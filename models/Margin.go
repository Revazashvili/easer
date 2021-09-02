package models

type Margin struct {
	Top    uint `json:"top" bson:"top"`
	Bottom uint `json:"bottom" bson:"bottom"`
	Left   uint `json:"left" bson:"left"`
	Right  uint `json:"right" bson:"right"`
}
