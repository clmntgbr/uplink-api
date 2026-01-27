# Executables (local)
DOCKER_COMP = docker compose

# Docker containers
PHP_CONT = $(DOCKER_COMP) exec php

# Misc
.DEFAULT_GOAL = help
.PHONY        : help up down remove sh trust-cert env entity schema install-hooks

## —— 🎵 🐳 The Symfony Docker Makefile 🐳 🎵 ——————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Docker 🐳 ————————————————————————————————————————————————————————————————
up: ## Start the docker hub in detached mode (no logs)
	@$(DOCKER_COMP) up -d --wait

down: ## Stop the docker hub
	@$(DOCKER_COMP) down --remove-orphans

remove: ## Remove all containers and volumes
	@$(DOCKER_COMP) down --remove-orphans -v

sh: ## Connect to the PHP container
	@$(PHP_CONT) sh

trust-cert: ## Install local SSL certificate
	@echo "Installing local SSL certificate..."
	@docker cp php:/data/caddy/pki/authorities/local/root.crt /tmp/root.crt
	@if [ "$$(uname)" = "Darwin" ]; then \
		echo "Detected macOS. Installing certificate..."; \
		sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain /tmp/root.crt; \
		echo "Certificate installed successfully!"; \
	elif [ "$$(uname)" = "Linux" ]; then \
		echo "Detected Linux. Installing certificate..."; \
		sudo cp /tmp/root.crt /usr/local/share/ca-certificates/root.crt; \
		sudo update-ca-certificates; \
		echo "Certificate installed successfully!"; \
	elif [ "$$(uname)" = "MINGW64_NT" ] || [ "$$(uname)" = "MINGW32_NT" ]; then \
		echo "Detected Windows. Opening certificate installer..."; \
		certutil -addstore -f "ROOT" /tmp/root.crt; \
		echo "Certificate installed successfully!"; \
	else \
		echo "Unknown operating system. Please install the certificate manually from: /tmp/root.crt"; \
	fi
	@rm /tmp/root.crt

env: ## Show environment variables
	@$(PHP_CONT) bin/console debug:dotenv

## —— Git 🔧 ———————————————————————————————————————————————————————————————————
install-hooks: ## Install git hooks for code quality
	@echo "Installing git hooks..."
	@cp .githooks/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "✅ Git hooks installed successfully!"

## —— Symfony 🎵 ———————————————————————————————————————————————————————————————
rector: ## Run rector
	@$(PHP_CONT) composer rector-fix

phpstan: ## Run phpstan
	@$(PHP_CONT) composer phpstan

php-cs-fixer: ## Run php-cs-fixer
	@$(PHP_CONT) composer cs-fix

qa: ## Run all qa tools
	@$(PHP_CONT) composer qa-fix