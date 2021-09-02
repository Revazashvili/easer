package models

import "time"

type Template struct {
	Id           string    `json:"id" json:"_id"`
	Name         string    `json:"name" bson:"name"`
	Description  string    `json:"description" bson:"description"`
	Owner        string    `json:"owner" bson:"owner"`
	TemplateBody string    `json:"templateBody" bson:"template_body"`
	Options      Options   `json:"options" bson:"options"`
	Created      time.Time `json:"created" bson:"created"`
	CreatedBy    string    `json:"created_by" bson:"created_by"`
	Updated      time.Time `json:"updated" bson:"updated"`
	UpdatedBy    string    `json:"updated_by" bson:"updated_by"`
}
