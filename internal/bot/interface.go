package bot

import (
	"context"

	"github.com/getflow/feedback-service/internal/bot/models"
)

type Bot interface {
	SendCommand(ctx context.Context, command *models.Command) error
}
