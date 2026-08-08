package bot

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	maxbot "github.com/max-messenger/max-bot-api-client-go/v2"
	maxmodel "github.com/max-messenger/max-bot-api-client-go/v2/model"

	"github.com/getflow/feedback-service/internal/bot/models"
)

type MaxBot struct {
	Bot *maxbot.Api
}

func (b MaxBot) SendCommand(ctx context.Context, command *models.Command) error {
	message := maxbot.NewMessage()
	chatID, err := strconv.ParseInt(command.ChatID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid chat ID: %w", err)
	}

	// TODO: Add group chats support
	message.SetUser(chatID)
	message.SetText(command.Text)
	message.SetFormat(maxmodel.FormatHTML)

	_, err = b.Bot.Messages.Send(ctx, message)
	return err
}

func NewMaxBot(token string) (*MaxBot, error) {
	var tlsConfig *tls.Config

	caCert := os.Getenv("CERT_PEM")
	if len(caCert) == 0 {
		return nil, fmt.Errorf("No certificate provided")
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM([]byte(caCert)) {
		return nil, fmt.Errorf("failed to add CA cert")
	}
	tlsConfig = &tls.Config{RootCAs: pool}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	opts := []maxbot.Opt{
		maxbot.WithHTTPClient(&http.Client{
			Timeout:   10 * time.Second,
			Transport: tr,
		}),
	}
	api, err := maxbot.NewApi(token, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to initialize bot %w", err)
	}
	return &MaxBot{
		Bot: api,
	}, nil
}
