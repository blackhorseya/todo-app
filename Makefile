app_name = todo
app_version = latest
project_id = sean-side-uat
namespace = sean-side-uat-ns

.PHONY: build-image
build-image:
	@docker build -t $(app_name):$(app_version) \
	--label "app.name=$(app_name)" \
	--label "app.version=$(app_version)" \
	--pull .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(app_name)

.PHONY: run-with-docker
run-with-docker:
	@docker run -it --rm -p 8080:8080 \
	-v $(shell pwd)/configs/app.yaml:/app/configs/app.yaml \
	$(app_name):$(app_version)

.PHONY: run-mongo
run-mongo:
	@helm --namespace $(namespace) upgrade --install $(app_name)-db bitnami/mongodb \
	--values ./deployments/mongo/local.yaml

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(app_name) -q`

.PHONY: tag-image
tag-image:
	@docker tag $(app_name):$(app_version) gcr.io/$(project_id)/$(app_name):$(app_version)

.PHONY: push-image
push-image:
	@docker push gcr.io/$(project_id)/$(app_name):$(app_version)

.PHONY: deploy-with-helm
deploy-with-helm:
	@helm --namespace $(namespace) \
	upgrade --install $(app_name) ./deployments/helm \
	--values ./deployments/helm/values.yaml

.PHONY: gen
gen: gen-pb gen-swagger gen-wire

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities \
	--proto_path=./internal/app/protos \
	./internal/app/protos/*.proto

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/app/main.go --parseInternal -o internal/app/docs

.PHONY: test-with-coverage
test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/todo-app'

.PHONY: install-mod
download-mod:
	@go mod download

.PHONY: install-tools
install-tools:
	@go get -v golang.org/x/lint/golint
