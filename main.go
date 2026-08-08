package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/getflow/feedback-service/internal/bot"
	"github.com/getflow/feedback-service/internal/bot/models"
)

type Feedback struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (feedback *Feedback) Format() string {
	value := reflect.ValueOf(feedback).Elem()
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).String() == "" {
			value.Field(i).SetString("-")
		}
	}

	return fmt.Sprintf(
		"<b>Отправитель:</b> %s\n<b>Компания:</b> %s\n<b>Телефон:</b> %s\n<b>Почта:</b> %s\n\n<b>Сообщение:</b> %s",
		feedback.Name, feedback.Company, feedback.Phone, feedback.Email, feedback.Message,
	)
}

func main() {
	b, err := bot.GetBot()
	if err != nil {
		log.Fatalf("failed creating bot api object: %s", err)
	}

	r := gin.Default()
	r.POST("/feedback", func(c *gin.Context) {
		var feedback *Feedback
		if err := json.NewDecoder(c.Request.Body).Decode(&feedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error while processing a request",
			})
			return
		}

		text := feedback.Format()

		if err = b.SendCommand(context.Background(), &models.Command{
			ChatID:    os.Getenv("FB_CHANNEL"),
			Text:      text,
		}); err != nil {
			log.Printf("error while sending internal request: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error while sending internal request",
			})
			return
		}

		c.JSON(http.StatusOK, feedback)
	})

	log.Printf("Starting application....")

	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("FB_PORT")))
}
