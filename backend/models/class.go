package models

type Class struct {
    ClassID int `json:"class_id" bson:"class_id"`
    Name    string `json:"name" bson:"name"`
}