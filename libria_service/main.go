package main

import (
	"database/sql"
	"fmt"
	"libria/answers"
	"libria/topics"
	"libria/utility"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

var dbClient *sql.DB

func startup() {
	dbClient = utility.NewDbClient()
	for utility.Migrate(dbClient) != nil {
		fmt.Printf("Verbindung Fehlgeschlagen, %s", utility.Migrate(dbClient))
		time.Sleep(20 * time.Second)
	}
}

func main() {
	startup()
	TopicRepository := topics.NewRepository(dbClient)
	TopicService := topics.NewService(TopicRepository)
	TopicDelivery := topics.NewDelivery(TopicService)

	AnswerRepository := answers.NewRepository(dbClient)
	AnswerService := answers.NewService(AnswerRepository)
	AnswerDelivery := answers.NewDelivery(AnswerService)

	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/topics", TopicDelivery.GetAll)
	e.GET("/api/topic/:id", TopicDelivery.GetById)
	e.GET("/api/randomTopic", TopicDelivery.GetRandom)
	e.POST("/api/topic", TopicDelivery.Post)
	e.PUT("/api/topic/:id", TopicDelivery.Update)
	e.DELETE("/api/topic/:id", TopicDelivery.Delete)

	e.GET("/api/answers", AnswerDelivery.GetAll)
	e.GET("/api/answers/:topic_id", AnswerDelivery.GetAllByTopic)
	e.GET("/api/answer/:id", AnswerDelivery.GetById)
	e.POST("/api/answer", AnswerDelivery.Post)
	e.PUT("/api/answer/:id", AnswerDelivery.Update)
	e.DELETE("/api/answer/:id", AnswerDelivery.Delete)

	e.Logger.Fatal(e.Start(":1324"))
}
