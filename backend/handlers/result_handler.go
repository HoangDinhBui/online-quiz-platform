package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
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

func (h *ResultHandler) GenerateUserID(c *gin.Context) {
    userID := uuid.New().String()
    ctx := context.Background()
    h.collection.Database().Client().Database("quiz_platform").Collection("sessions").
        InsertOne(ctx, bson.M{"user_id": userID, "created_at": time.Now()})
    c.JSON(200, gin.H{"user_id": userID})
}

func (h *ResultHandler) SubmitAnswers(c *gin.Context) {
    var submission struct {
        UserID  string            `json:"user_id"`
        ClassID int            `json:"class_id"`
        Answers map[int]string `json:"answers"`
    }
    if err := c.ShouldBindJSON(&submission); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()
    questionsCursor, err := h.collection.Database().Collection("questions").Find(
        ctx,
        bson.M{"class_id": submission.ClassID},
    )
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    var questions []models.Question
    if err := questionsCursor.All(ctx, &questions); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    score := 0
    for _, q := range questions {
        if submission.Answers[q.QuestionID] == q.CorrectAnswer {
            score += 10
        }
    }

    result := models.Result{
        UserID:    submission.UserID,
        ClassID:   submission.ClassID,
        Score:     score,
        Timestamp: time.Now().String(),
    }
    _, err = h.collection.InsertOne(ctx, result)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"score": score})
}