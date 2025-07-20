package models

type Class struct {
    ClassID string `json:"class_id" bson:"class_id"`
    Name    string `json:"name" bson:"name"`
}