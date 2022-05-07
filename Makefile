.PHONY: test
test:
	@echo "\n🛠️  Running unit tests..."
	go test ./...

.PHONY: build
build:
	@echo "\n🔧  Building Go binaries..."
	GOOS=darwin GOARCH=amd64 go build -o bin/sql-prometheus-metrics-darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -o bin/sql-prometheus-metrics-linux-amd64 .

.PHONY: run
run: build
	@echo "\n🔧  Building Go binaries..."
	source .env && bin/sql-prometheus-metrics-darwin-amd64


.PHONY: docker-build
docker-build:
	@echo "\n📦 Building sql-prometheus-metrics Docker image..."
	docker build -t sql-prometheus-metrics:1.0 . 

# From this point `kind` is required
.PHONY: cluster
cluster:
	@echo "\n🔧 Creating Kubernetes cluster..."
	kind create cluster --config dev/manifests/kind/kind.cluster.yaml

.PHONY: delete-cluster
delete-cluster:
	@echo "\n♻️  Deleting Kubernetes cluster..."
	kind delete cluster

.PHONY: push
push: docker-build
	@echo "\n📦 Pushing admission-webhook image into Kind's Docker daemon..."
	kind load docker-image sql-prometheus-metrics:1.0

.PHONY: deploy
deploy: push delete-app app
	@echo "\n🚀 Deploying sql-prometheus-metrics..."
	kubectl apply -f dev/manifests/apps/
    kubectl -n apps logs -l app=sql-prometheus-metrics -f

.PHONY: app
app:
	@echo "\n🚀 Deploying sql-prometheus-metrics..."
	kubectl apply -f dev/manifests/apps/app.yaml
    
.PHONY: delete-app
delete-app:
	@echo "\n♻️ Deleting sql-prometheus-metrics..."
	kubectl delete -f dev/manifests/apps/app.yaml || true

.PHONY: db
db:
	@echo "\n🚀 Deploying postgres..."
	helm upgrade  --install postgres  bitnami/postgresql  --namespace apps --version 11.1.25 -f  dev/manifests/postgres/values.yaml 

.PHONY: logs
logs:
	@echo "\n🔍 Streaming sql-prometheus-metrics logs..."
	kubectl -n apps logs -l app=sql-prometheus-metrics -f

.PHONY: delete-all
delete-all: delete delete-config delete-pod delete-bad-pod
