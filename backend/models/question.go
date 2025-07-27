package models

type Question struct {
    QuestionID    int   `json:"question_id" bson:"question_id"`
    ClassID       int   `json:"class_id" bson:"class_id"`
    Content       string   `json:"content" bson:"content"`
    Options       []string `json:"options" bson:"options"`
    CorrectAnswer string   `json:"correct_answer" bson:"correct_answer"`
}