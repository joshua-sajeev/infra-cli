.PHONY: help test-env-up test-env-up-detach test-env-restart test-env-logs test-env-down build deploy

help:
	@echo "Available commands:"
	@echo ""
	@echo "  make test-env-up         - Start all test servers"
	@echo "  make test-env-up-detach  - Start all servers in background"
	@echo "  make test-env-restart    - Restart all services"
	@echo "  make test-env-logs       - View logs from all services"
	@echo "  make test-env-down       - Stop all services"
	@echo "  make build               - Build the CLI"
	@echo "  make deploy              - Run deploy command"

test-env-up:
	cd infra-lab && docker compose up

test-env-up-detach:
	cd infra-lab && docker compose up -d

test-env-restart:
	cd infra-lab && docker compose restart

test-env-logs:
	cd infra-lab && docker compose logs -f

test-env-down:
	cd infra-lab && docker compose down

build:
	go build -o infra-cli

deploy:
	go run ./cmd deploy
