package main

import (
    "context"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/gorilla/websocket"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
    "github.com/HoangDinhBui/online-quiz-platform/backend/handlers"
    "os"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        conn.WriteMessage(websocket.TextMessage, msg)
    }
}

func main() {
    // Kết nối MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongodb:27017"))
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("online-quiz-platform")

    // Kết nối Redis
    redisClient := redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_URL"),
    })
    if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
        log.Fatal(err)
    }

    // Khởi tạo router Gin
    r := gin.Default()

    // Thêm middleware CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // Khởi tạo handlers
    classHandler := handlers.NewClassHandler(db, redisClient)
    questionHandler := handlers.NewQuestionHandler(db, redisClient)
    resultHandler := handlers.NewResultHandler(db)

    // Định nghĩa API
    r.GET("/classes", classHandler.GetClasses)
    r.GET("/questions", questionHandler.GetQuestions)
    r.POST("/submit", resultHandler.SubmitAnswers)
    r.GET("/generate-user-id", resultHandler.GenerateUserID)
    r.GET("/ws", handleWebSocket)
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // Chạy server
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
