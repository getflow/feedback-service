package bot

import (
	"context"
	"fmt"

	tgbot "github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"

	"github.com/getflow/feedback-service/internal/bot/models"
)


type TgBot struct {
	Bot *tgbot.Bot
}

func (b TgBot) SendCommand(ctx context.Context, command *models.Command) error {
	_, err := b.Bot.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID:    command.ChatID,
		Text:      command.Text,
		ParseMode: tgmodels.ParseModeHTML,
	})
	return err
}

func NewTgBot(token string) (*TgBot, error) {
	bot, err := tgbot.New(token)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize bot %w", err)
	}
	return &TgBot{
		Bot: bot,
	}, nil
}