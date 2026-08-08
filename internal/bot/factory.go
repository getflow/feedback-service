package bot

import (
	"fmt"
	"os"
	"strings"
)


const (
	BotTg = "BOT_TG"
	BotMax = "BOT_MAX"
)


func GetBot() (Bot, error) {
	token := os.Getenv("FB_TOKEN")
	if len(token)==0 {
		return nil, fmt.Errorf("No token provided")
	}
	botType := os.Getenv("FB_BOT_TYPE")
	switch strings.ToUpper(botType) {
	case BotTg:
		bot, err := NewTgBot(token)
		return bot, err
	case BotMax:
		bot, err := NewMaxBot(token)
		return bot, err
	default:
		return nil, fmt.Errorf("%s is not supported", botType)
	}
}