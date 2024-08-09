# Install dependencies
install: install_api install_ui
install_api:
	cd apps/api && \
	go mod download 

install_ui:
	cd apps/ui && \
	npm ci

# Run Dev Server
# make -j2 server_dev
server_dev: server_dev_api server_dev_ui
server_dev_api:
	cd apps/api && \
	CompileDaemon -command="./api"

server_dev_ui:
	cd apps/ui && \
	npm run dev

# Build
# make -j2 build
build: build_api build_ui
build_api:
	cd apps/api && \
	go build .

build_ui:
	cd apps/ui && \
	npm run build

# Run Prod Server
# make -j2 server_prod
server_prod: server_prod_api server_prod_ui
server_prod_api:
	cd apps/api && \
	./api

# NOTE: On prod, use nginx to serve static, not available without docker
# Still use dev for local development
server_prod_ui:
	cd apps/ui && \
	npm run dev

