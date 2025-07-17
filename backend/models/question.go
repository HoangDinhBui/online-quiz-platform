package models

type Question struct {
    QuestionID    string   `json:"question_id" bson:"question_id"`
    ClassID       string   `json:"class_id" bson:"class_id"`
    Content       string   `json:"content" bson:"content"`
    Options       []string `json:"options" bson:"options"`
    CorrectAnswer string   `json:"correct_answer" bson:"correct_answer"`
}