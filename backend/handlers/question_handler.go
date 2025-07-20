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

   type QuestionHandler struct {
       collection   *mongo.Collection
       redisClient  *redis.Client
   }

   func NewQuestionHandler(db *mongo.Database, redisClient *redis.Client) *QuestionHandler {
       return &QuestionHandler{
           collection:  db.Collection("questions"),
           redisClient: redisClient,
       }
   }

   func (h *QuestionHandler) GetQuestions(c *gin.Context) {
       classID := c.Query("class_id")
       cacheKey := "questions:" + classID
       ctx := context.Background()

       // Thử lấy từ Redis
       questionsData, err := h.redisClient.Get(ctx, cacheKey).Result()
       if err == redis.Nil {
           // Nếu không có trong Redis, lấy từ MongoDB
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

           // Lưu vào Redis với TTL 1 giờ
           questionsBytes, _ := json.Marshal(questions)
           h.redisClient.Set(ctx, cacheKey, questionsBytes, time.Hour)
           c.JSON(200, questions)
           return
       } else if err != nil {
           c.JSON(500, gin.H{"error": err.Error()})
           return
       }

       // Trả dữ liệu từ Redis
       var questions []models.Question
       json.Unmarshal([]byte(questionsData), &questions)
       c.JSON(200, questions)
   }