# Spider 抖音 API 服务

一个使用 Go + Gin 的轻量服务，用于抓取与代理抖音相关数据，包括：用户信息、用户搜索、用户作品列表、作品详情、评论与评论回复。

## 快速开始

- 环境要求
  - Go 1.20+（推荐）
  - Windows/macOS/Linux 均可
  - 开发环境建议安装 `swag` CLI 以生成/更新 Swagger 文档
- 安装依赖
  - 在项目根目录执行：`go mod download` 或 `go mod tidy`
- 运行服务
  - 创建并配置 `config.yaml`（见下文“配置说明”）
  - 执行：`go run main.go`
  - 或在 Windows 执行批处理：` .\.bat `（会构建 `spider.exe`）
- 服务默认端口
  - 由 `config.yaml` 的 `system.port` 指定，示例为 `8004`
  - 启动后可访问：`http://localhost:8004/`
  - Swagger 文档：`http://localhost:8004/swagger/index.html#/`

## 配置说明（config.yaml）

项目使用 `viper` 读取根目录下的 `config.yaml`。`.gitignore` 默认忽略该文件，请自行创建。

最小可运行示例（请根据实际环境修改）：

```yaml
cookie:
  douyin: "ttwid=...; msToken=...; odin_tt=...; __ac_nonce=...; __ac_signature=...; ..."

system:
  name: "spider-run"
  version: "1.0.1"
  port: "8004"
  env: "dev"   # dev/test/release

logging:
  level: "info"
  format: "text"
  rotation_time: "24h"
  max_age: "60d"
  output:
    - type: "console"
      colored: true
  categories:
    access:
      enabled: true
      path: "./cache/logs/info"
    slow:
      enabled: true
      threshold: 1
      path: "./cache/logs/slow"
    error:
      enabled: true
      path: "./cache/logs/err"

# 可选：Redis 配置（如不需要可删除或保持默认）
redis:
  - host: "127.0.0.1"
    port: 6379
    password: "123456"
    db: 0
    dbName: "default"
    maxIdle: 10
    maxActive: 20
```

重要说明：
- `cookie.douyin` 必填。抖音 Web API 强依赖 Cookie（含 `ttwid`、`msToken`、`__ac_nonce`、`__ac_signature` 等）。请从浏览器登录抖音 Web 后复制 Cookie，并保持更新。
- `system.env` 为 `dev` 时，程序启动会自动尝试运行 `swag init` 以生成文档，需要本机安装 `swag` 并在环境变量 `PATH` 中可用。

安装 `swag` CLI：

```bash
go install github.com/swaggo/swag/cmd/swag@latest
# 请确保 GOPATH/bin 在 PATH 中，例如 Windows 下：%USERPROFILE%\go\bin
```

## 路由与中间件

- 路由前缀：`/api/douyin`
- 中间件：
  - 请求 ID（`RequestID`）
  - 访问日志（`AccessLogMiddleware`）
  - 慢日志（`SlowLogMiddleware`，阈值来自配置）
  - CORS（`CorsMiddleware`）
- Swagger UI：`/swagger/index.html#/`

## API 列表与示例

以下接口均为 GET 请求，返回 JSON。示例默认端口为 `8004`，请按需替换。

### 用户相关

- 用户信息
  - 路径：`/api/douyin/user/info`
  - 参数：`secId`（string，必填）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/user/info?secId=MS4wLjABAAAA..."
    ```

- 用户搜索
  - 路径：`/api/douyin/user/search`
  - 参数：`keyword`（string，必填）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/user/search?keyword=美食"
    ```

- 用户作品列表
  - 路径：`/api/douyin/user/video`
  - 参数：`secId`（string，必填）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/user/video?secId=MS4wLjABAAAA..."
    ```

### 作品相关

- 作品详情
  - 路径：`/api/douyin/aweme/detail`
  - 参数：`awemeId`（string，必填）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/aweme/detail?awemeId=7561795410549853481"
    ```

### 评论相关

- 评论列表
  - 路径：`/api/douyin/comment/list`
  - 参数：`awemeId`（string，必填）、`cursor`（string，默认 `0`）、`count`（string，默认 `10`）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/comment/list?awemeId=7561795410549853481&cursor=0&count=10"
    ```

- 评论回复列表
  - 路径：`/api/douyin/comment/reply`
  - 参数：`itemId`（string，必填）、`commentId`（string，必填）、`cursor`（string，默认 `0`）、`count`（string，默认 `10`）
  - 示例：
    ```bash
    curl "http://localhost:8004/api/douyin/comment/reply?itemId=7561795410549853481&commentId=7565069968577479451&cursor=0&count=10"
    ```

## 目录与代码位置

- 入口：`main.go`
- 路由：`internal/router/router.go`、`internal/router/douyin/douyin.go`
- 处理器：`internal/app/handler/douyin/http_user.go`、`http_aweme.go`、`http_comment.go`
- 业务逻辑：`internal/app/logic/douyin/*`
- 配置：`internal/config/*.go`、`config.yaml`
- Swagger 生成：`cmd/swag_cmd.go`（开发环境自动执行 `swag init`）

## 常见问题

- 403/签名校验失败
  - 检查并更新 `cookie.douyin`，确保包含关键字段（`ttwid`、`msToken`、`__ac_nonce`、`__ac_signature` 等）。
  - 保持合适的 `User-Agent`（代码已内置，亦可在逻辑层调整）。
- Swagger 页面空白或 404
  - 确认 `system.env` 为 `dev`，且本机安装了 `swag` 并在 PATH 中可执行。
  - 或手动在项目根目录执行 `swag init` 重新生成 `docs`。
- 端口冲突
  - 修改 `config.yaml` 中的 `system.port`，重新启动。

## 许可证

本项目为学习与技术研究用途，接口数据来源于抖音 Web。使用时请遵守相关平台的服务条款与法律法规。

## 平台目录表格

为后续扩展其他平台（如快手、微博、小红书等），约定统一的目录结构与路由前缀，并在下表中进行区分：

| 平台 | 路由前缀 | 路由注册 | Handler 目录 | Logic 目录 | 类型定义 |
|------|-----------|-----------|--------------|------------|----------|
| Douyin | `/api/douyin` | `internal/router/douyin/douyin.go#InitDouyinRoute` | `internal/app/handler/douyin/*`（`http_user.go`、`http_aweme.go`、`http_comment.go`） | `internal/app/logic/douyin/*` | `internal/app/types/types_douyin/*` |
| Kuaishou（预留） | `/api/kuaishou` | `internal/router/kuaishou/kuaishou.go` | `internal/app/handler/kuaishou/*` | `internal/app/logic/kuaishou/*` | `internal/app/types/types_kuaishou/*` |
| XHS（预留） | `/api/xhs` | `internal/router/xhs/xhs.go` | `internal/app/handler/xhs/*` | `internal/app/logic/xhs/*` | `internal/app/types/types_xhs/*` |
| Weibo（预留） | `/api/weibo` | `internal/router/weibo/weibo.go` | `internal/app/handler/weibo/*` | `internal/app/logic/weibo/*` | `internal/app/types/types_weibo/*` |

### 新增平台约定与步骤

- 路由前缀统一采用：`/api/<platform>`，例如 `/api/kuaishou`。
- 在 `internal/router/<platform>/<platform>.go` 中实现 `Init<Platform>Route(r *gin.Engine)`，并在 `router.InitRouter(r)` 中调用。
- 在 `internal/app/handler/<platform>` 下新增 `http_*.go`，按功能拆分（如 `http_user.go`、`http_aweme.go`、`http_comment.go`）。
- 在 `internal/app/logic/<platform>` 实现对应客户端与方法（例如 `<Platform>Client` 与具体接口方法）。
- 在 `internal/app/types/types_<platform>` 定义该平台的请求/响应结构体，命名与 Douyin 保持一致性（如 `UserInfoReq`、`CommentListReq` 等）。
- 如需生成 Swagger 文档，补充注释与标签后执行 `swag init`。