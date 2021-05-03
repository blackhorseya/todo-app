APP_NAME = todo
VERSION = latest
PROJECT_ID = sean-side
NS = side
DEPLOY_TO = uat

DB_URI="mongodb://todo-app:changeme@localhost:27017/todo-db"

.PHONY: clean
clean:
	@rm -rf coverage.txt profile.out bin

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@golint ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/todo-app'

.PHONY: build-image
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--pull .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: tag-image
tag-image:
	@docker tag $(APP_NAME):$(VERSION) gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: push-image
push-image:
	@docker push gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: install-db
install-db:
	@helm --namespace $(NS) upgrade --install $(APP_NAME)-db bitnami/mongodb \
	--values ./deployments/configs/$(DEPLOY_TO)/mongo.yaml

.PHONY: deploy
deploy:
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) ./deployments/helm \
	--values ./deployments/configs/$(DEPLOY_TO)/todo.yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-pb gen-swagger gen-wire

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/pkg/entity ./internal/pkg/entity/**/*.proto

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs

.PHONY: migrate-up
migrate-up:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down