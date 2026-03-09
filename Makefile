.PHONY: help up up-detach restart logs down 

help:
	@echo "Available commands:"
	@echo ""
	@echo "  make up                  - Start all test servers"
	@echo "  make up-detach           - Start all servers in background"
	@echo "  make restart             - Restart all services"
	@echo "  make logs                - View logs from all services"
	@echo "  make down                - Stop all services"

up:
	docker-compose up

up-detach:
	docker-compose up -d

restart:
	docker-compose restart

logs:
	docker-compose logs -f

down:
	docker-compose down
