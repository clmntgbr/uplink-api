.PHONY: dev dev-logs dev-down dev-restart dev-rebuild prod prod-logs prod-down prod-restart build clean shell test help

# ============================================
# Development commands (docker-compose.yml)
# ============================================

dev:
	@echo "🚀 Starting development environment in background..."
	docker-compose up -d --build

dev-logs:
	@echo "📋 Showing development logs..."
	docker-compose logs -f

dev-down:
	@echo "🛑 Stopping development environment..."
	docker-compose down

dev-restart:
	@echo "🔄 Restarting development environment..."
	docker-compose restart

dev-rebuild:
	@echo "🔨 Rebuilding development environment..."
	docker-compose down
	docker-compose up --build

# ============================================
# Production commands (docker-compose.prod.yml)
# ============================================

prod:
	@echo "🚀 Starting production environment..."
	docker-compose -f docker-compose.prod.yml up --build

prod-d:
	@echo "🚀 Starting production environment in background..."
	docker-compose -f docker-compose.prod.yml up -d --build

prod-logs:
	@echo "📋 Showing production logs..."
	docker-compose -f docker-compose.prod.yml logs -f

prod-down:
	@echo "🛑 Stopping production environment..."
	docker-compose -f docker-compose.prod.yml down

prod-restart:
	@echo "🔄 Restarting production environment..."
	docker-compose -f docker-compose.prod.yml restart

prod-rebuild:
	@echo "🔨 Rebuilding production environment..."
	docker-compose -f docker-compose.prod.yml down
	docker-compose -f docker-compose.prod.yml up --build

# ============================================
# Build specific images
# ============================================

build-dev:
	@echo "🔨 Building development image..."
	docker build --target development -t uplink-api:dev .

build-prod:
	@echo "🔨 Building production image..."
	docker build --target production -t uplink-api:prod .

# ============================================
# Utility commands
# ============================================

shell:
	@echo "🐚 Opening shell in development container..."
	docker-compose exec api sh

shell-prod:
	@echo "🐚 Opening shell in production container..."
	docker-compose -f docker-compose.prod.yml exec api sh

test:
	@echo "🧪 Running tests in development container..."
	docker-compose exec api go test ./... -v

clean:
	@echo "🧹 Cleaning up..."
	docker-compose down -v
	docker-compose -f docker-compose.prod.yml down -v
	rm -rf tmp/
	docker system prune -f

clean-all:
	@echo "🧹 Deep cleaning (including images)..."
	docker-compose down -v --rmi all
	docker-compose -f docker-compose.prod.yml down -v --rmi all
	rm -rf tmp/
	docker system prune -af --volumes

# Show image sizes
size:
	@echo "📊 Image sizes:"
	@echo "\n🔧 Development:"
	@docker images | grep uplink-api | grep dev || echo "Dev image not built"
	@echo "\n🚀 Production:"
	@docker images | grep uplink-api | grep prod || echo "Prod image not built"

# Health checks
health:
	@echo "🏥 Health check:"
	@echo "\n🔧 Development:"
	@curl -f http://localhost:3000/ 2>/dev/null && echo "✅ Dev OK" || echo "❌ Dev not running"
	@echo "\n🚀 Production:"
	@curl -f http://localhost:3000/ 2>/dev/null && echo "✅ Prod OK" || echo "❌ Prod not running"

# Show running containers
ps:
	@echo "📦 Running containers:"
	@docker-compose ps
	@docker-compose -f docker-compose.prod.yml ps

# Show logs from both environments
logs-all:
	@echo "📋 All logs:"
	docker-compose logs
	docker-compose -f docker-compose.prod.yml logs

# ============================================
# Help
# ============================================

help:
	@echo "🎯 Available commands:"
	@echo ""
	@echo "Development (docker-compose.yml):"
	@echo "  make dev              - Start dev with hot reload"
	@echo "  make dev-d            - Start dev in background"
	@echo "  make dev-logs         - Show dev logs"
	@echo "  make dev-down         - Stop dev"
	@echo "  make dev-restart      - Restart dev"
	@echo "  make dev-rebuild      - Rebuild and restart dev"
	@echo ""
	@echo "Production (docker-compose.prod.yml):"
	@echo "  make prod             - Start production"
	@echo "  make prod-d           - Start production in background"
	@echo "  make prod-logs        - Show production logs"
	@echo "  make prod-down        - Stop production"
	@echo "  make prod-restart     - Restart production"
	@echo "  make prod-rebuild     - Rebuild and restart production"
	@echo ""
	@echo "Build:"
	@echo "  make build-dev        - Build dev image only"
	@echo "  make build-prod       - Build prod image only"
	@echo ""
	@echo "Utils:"
	@echo "  make shell            - Shell in dev container"
	@echo "  make shell-prod       - Shell in prod container"
	@echo "  make test             - Run tests in dev"
	@echo "  make clean            - Clean containers and volumes"
	@echo "  make clean-all        - Deep clean (including images)"
	@echo "  make size             - Show image sizes"
	@echo "  make health           - Health check"
	@echo "  make ps               - Show running containers"
	@echo "  make help             - Show this help"