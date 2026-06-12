SERVICE_NAME = my-service
API_DIR = /Users/zhaocong/Documents/space/gospace/mt_node
SQL_FILE_DIR = /Users/zhaocong/Documents/space/gospace/mt_node/files/sql
PG_URL= postgres://postgres:postgres@127.0.0.1:5432/mt_node?sslmode=disable
OUTPUT_DIR = output

# 默认目标
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "GoZero 项目构建工具"
	@echo ""
	@echo "命令列表:"
	@echo "  make gen      - 生成代码 (api/model/rpc)"
	@echo "  make build    - 编译项目"
	@echo "  make run      - 运行服务"
	@echo "  make test     - 运行测试"
	@echo "  make clean    - 清理文件"
	@echo "  make docker   - 构建 Docker 镜像"

# 代码生成
.PHONY: gen
gen:
	@echo "正在生成代码..."
	# 生成 API
	goctl api go -api $(API_DIR)/*.api -dir .
	# 生成 model
	# goctl model pg datasource --url "$(PG_URL)" --schema "mt_auth" --table "chain_info" --dir "./model" -c
	@echo "代码生成完成!"

# 安装依赖
.PHONY: deps
deps:
	@echo "正在安装依赖..."
	go mod tidy
	@echo "依赖安装完成!"

# 编译
.PHONY: build
build: deps
	@echo "正在编译..."
	mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(SERVICE_NAME) cmd/$(SERVICE_NAME)/main.go
	@echo "编译完成: $(OUTPUT_DIR)/$(SERVICE_NAME)"

# 运行
.PHONY: run
run: build
	@echo "启动服务..."
	./$(OUTPUT_DIR)/$(SERVICE_NAME)

# 测试
.PHONY: test
test:
	@echo "运行测试..."
	go test -v ./...

# 清理
.PHONY: clean
clean:
	@echo "清理文件..."
	rm -rf $(OUTPUT_DIR)
	rm -f *.log
	@echo "清理完成!"

# Docker
.PHONY: docker
docker: build
	@echo "构建 Docker 镜像..."
	docker build -t $(SERVICE_NAME):latest .
	@echo "Docker 镜像构建完成!"