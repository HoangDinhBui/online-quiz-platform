package handlers

import (
    "context"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/HoangDinhBui/online-quiz-platform/backend/models"
    "time"
)

type ClassHandler struct {
    collection *mongo.Collection
    redisClient  *redis.Client
}

func NewClassHandler(db *mongo.Database, redisClient *redis.Client) *ClassHandler {
    return &ClassHandler{
           collection:  db.Collection("classes"),
           redisClient: redisClient,
       }
}

func (h *ClassHandler) GetClasses(c *gin.Context) {
       ctx := context.Background()

       // Thử lấy từ Redis
       classesData, err := h.redisClient.Get(ctx, "classes").Result()
       if err == redis.Nil {
           // Nếu không có trong Redis, lấy từ MongoDB
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

           // Lưu vào Redis với TTL 1 giờ
           classesBytes, _ := json.Marshal(classes)
           h.redisClient.Set(ctx, "classes", classesBytes, time.Hour)
           c.JSON(200, classes)
           return
       } else if err != nil {
           c.JSON(500, gin.H{"error": err.Error()})
           return
       }

       // Trả dữ liệu từ Redis
       var classes []models.Class
       json.Unmarshal([]byte(classesData), &classes)
       c.JSON(200, classes)
   }