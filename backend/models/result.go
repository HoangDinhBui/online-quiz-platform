package models

type Result struct {
    UserID   string `json:"user_id" bson:"user_id"`
    ClassID  string `json:"class_id" bson:"class_id"`
    Score    int    `json:"score" bson:"score"`
    Timestamp string `json:"timestamp" bson:"timestamp"`
}