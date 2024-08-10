# Install dependencies
install: install_api install_leaderboard install_ui
install_api:
	cd apps/api && \
	go mod download 
install_leaderboard:
	cd apps/leaderboard && \
	go mod download 
install_ui:
	cd apps/ui && \
	npm ci

# Run Dev Server
# make -j3 server_dev
server_dev: server_dev_api server_dev_leaderboard server_dev_ui
server_dev_api:
	cd apps/api && \
	CompileDaemon -command="./api"
server_dev_leaderboard:
	cd apps/leaderboard && \
	CompileDaemon -command="./leaderboard"
server_dev_ui:
	cd apps/ui && \
	npm run dev

# Build
# make -j3 build
build: build_api build_leaderboard build_ui
build_api:
	cd apps/api && \
	go build .
build_leaderboard:
	cd apps/leaderboard && \
	go build .
build_ui:
	cd apps/ui && \
	npm run build

# Run Prod Server
# make -j3 server_prod
server_prod: server_prod_api server_prod_leaderboard server_prod_ui
server_prod_api:
	cd apps/api && \
	./api
server_prod_leaderboard:
	cd apps/leaderboard && \
	./leaderboard

# NOTE: On prod, use nginx to serve static, not available without docker
# Still use dev for local development
server_prod_ui:
	cd apps/ui && \
	npm run dev

