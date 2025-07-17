package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/HoangDinhBui/online-quiz-platform/backend/models"
)

type ClassHandler struct {
    collection *mongo.Collection
}

func NewClassHandler(db *mongo.Database) *ClassHandler {
    return &ClassHandler{collection: db.Collection("classes")}
}

func (h *ClassHandler) GetClasses(c *gin.Context) {
    ctx := context.Background()
    cursor, err := h.collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    var classes []models.Class
    if err := cursor.All(ctx, &classes); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, classes)
}