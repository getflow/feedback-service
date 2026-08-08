# Feedback Service

Service to forward feedback messages to Telegram or MAX Messenger.

[![License: GPL v2](https://img.shields.io/badge/License-GPL_v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html)
![Go](https://img.shields.io/badge/Go-1.24-brightgreen)
![Docker](https://img.shields.io/badge/Docker-ready-blue)

## Overview

`feedback-service` is a lightweight Go HTTP service that receives feedback submissions via a REST API and forwards them as formatted messages to a bot channel on **Telegram** or **MAX Messenger**. It is designed to be embedded into any web application as a backend for a contact/feedback form.

### Features

- **Simple REST API** — one `POST /feedback` endpoint accepting JSON payloads
- **Multi-platform bot support** — switch between Telegram and MAX Messenger via environment variable
- **Docker-ready** — multi-stage build with UPX compression, ready for containerized deployment
- **CI/CD integrated** — GitLab CI pipeline for automated build and Docker image publishing

## Project Structure

```
.
├── main.go                  # Entry point, HTTP server, Feedback struct
├── internal/
│   └── bot/
│       ├── factory.go       # Bot factory (Telegram / MAX)
│       ├── interface.go     # Bot interface definition
│       ├── tg.go            # Telegram bot implementation
│       ├── max.go           # MAX Messenger bot implementation
│       └── models/
│           └── command.go   # Command model (ChatID + Text)
├── Dockerfile
├── .gitlab-ci.yml
├── go.mod
└── go.sum
```

## Quick Start

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) (for containerized deployment)

### Running Locally

```bash
# Set required environment variables
export FB_TOKEN="your-bot-token"
export FB_BOT_TYPE="BOT_TG"          # or "BOT_MAX"
export FB_CHANNEL="123456789"    # target chat/channel ID
export FB_PORT="3000"

# For MAX Messenger, also set:
export CERT_PEM="$(cat path/to/cert.pem)"

# Start the server
go run main.go
```

### Building with Docker

```bash
docker build -t getflow/feedback-service:latest .
docker run -p 3000:3000 \
  -e FB_TOKEN=your-bot-token \
  -e FB_BOT_TYPE=TG \
  -e FB_CHANNEL=123456789 \
  getflow/feedback-service:latest
```

## API Reference

### POST /feedback

Submits a feedback message. The service formats the data and sends it to the configured bot channel.

**Request Body** (`application/json`):

```json
{
  "name": "John Doe",
  "company": "Acme Corp",
  "phone": "+1234567890",
  "email": "john@example.com",
  "message": "I have a question about your service."
}
```

| Field     | Type   | Required | Description              |
|-----------|--------|----------|--------------------------|
| `name`    | string | Yes      | Sender's name       |
| `company` | string | No       | Sender's company         |
| `phone`   | string | No       | Sender's phone number    |
| `email`   | string | No       | Sender's email address   |
| `message` | string | Yes      | Feedback message body    |

**Responses:**

| Status Code | Description                                    |
|-------------|------------------------------------------------|
| `200 OK`    | Feedback accepted and forwarded successfully   |
| `400 Bad Request` | Invalid JSON payload                     |
| `500 Internal Server Error` | Failed to send message to bot channel |

**Example (curl):**

```bash
curl -X POST http://localhost:3000/feedback \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "company": "Acme Corp",
    "phone": "+1234567890",
    "email": "john@example.com",
    "message": "I love your product!"
  }'
```

## Environment Variables

| Variable      | Required | Description                                    |
|---------------|----------|------------------------------------------------|
| `FB_TOKEN`    | Yes      | Bot API token (Telegram or MAX)                |
| `FB_BOT_TYPE` | Yes      | Bot platform: `BOT_TG` (Telegram) or `BOT_MAX`         |
| `FB_CHANNEL`  | Yes      | Target chat / channel ID for message delivery  |
| `FB_PORT`     | No       | HTTP server port (default: `3000`)             |
| `CERT_PEM`    | Conditional | CA certificate in PEM format (required for MAX) |

## CI/CD

This project uses GitLab CI. The pipeline includes:

1. **docs** — Publishes the README to Docker Hub
2. **build** — Builds the Docker image and pushes it to Docker Hub

## License

This project is licensed under the [GNU General Public License v2.0](LICENSE).

## Author

**Getflow Tech** — [github.com/getflow](https://github.com/getflow)
