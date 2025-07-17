package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/HoangDinhBui/online-quiz-platform/backend/models"
    "time"
)

type ResultHandler struct {
    collection *mongo.Collection
}

func NewResultHandler(db *mongo.Database) *ResultHandler {
    return &ResultHandler{collection: db.Collection("results")}
}

func (h *ResultHandler) SubmitAnswers(c *gin.Context) {
    var submission struct {
        UserID   string            `json:"user_id"`
        ClassID  string            `json:"class_id"`
        Answers  map[string]string `json:"answers"` // question_id: answer
    }
    if err := c.ShouldBindJSON(&submission); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Giả sử chấm điểm bằng cách so sánh với correct_answer trong DB
    questionsCursor, err := h.collection.Database().Collection("questions").Find(
        context.Background(),
        bson.M{"class_id": submission.ClassID},
    )
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    var questions []models.Question
    if err := questionsCursor.All(context.Background(), &questions); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    score := 0
    for _, q := range questions {
        if submission.Answers[q.QuestionID] == q.CorrectAnswer {
            score += 10 // 10 điểm mỗi câu
        }
    }

    // Lưu kết quả
    result := models.Result{
        UserID:    submission.UserID,
        ClassID:   submission.ClassID,
        Score:     score,
        Timestamp: time.Now().String(),
    }
    _, err = h.collection.InsertOne(context.Background(), result)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"score": score})
}