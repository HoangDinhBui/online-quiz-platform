package main

import (
    "context"
    "log"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/HoangDinhBui/online-quiz-platform/backend/handlers"
    "time"
)

func main() {
    // Kết nối MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("quiz_platform")

    // Khởi tạo router Gin
    r := gin.Default()

    // Thêm cấu hình CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Khởi tạo handlers
    classHandler := handlers.NewClassHandler(db)
    questionHandler := handlers.NewQuestionHandler(db)
    resultHandler := handlers.NewResultHandler(db)

    // Định nghĩa API
    r.GET("/classes", classHandler.GetClasses)
    r.GET("/questions", questionHandler.GetQuestions)
    r.POST("/submit", resultHandler.SubmitAnswers)

    // Chạy server
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
