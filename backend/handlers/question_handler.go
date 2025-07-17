package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/HoangDinhBui/online-quiz-platform/backend/models"
)

type QuestionHandler struct {
    collection *mongo.Collection
}

func NewQuestionHandler(db *mongo.Database) *QuestionHandler {
    return &QuestionHandler{collection: db.Collection("questions")}
}

func (h *QuestionHandler) GetQuestions(c *gin.Context) {
    classID := c.Query("class_id")
    ctx := context.Background()
    cursor, err := h.collection.Find(ctx, bson.M{"class_id": classID})
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    var questions []models.Question
    if err := cursor.All(ctx, &questions); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, questions)
}