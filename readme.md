# aweme

一个轻量级短视频应用程序后端

[![Codacy Badge](https://img.shields.io/codacy/grade/581a7a1d9685419eb6fef5afb636544d?style=for-the-badge)](https://www.codacy.com/gh/PlanVX/aweme/dashboard?utm_source=github.com&utm_medium=referral&utm_content=PlanVX/aweme&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/PlanVX/aweme?style=for-the-badge)](https://goreportcard.com/report/github.com/PlanVX/aweme)
[![build](https://img.shields.io/github/actions/workflow/status/PlanVx/aweme/ci.yaml?style=for-the-badge)](https://github.com/PlanVX/aweme/actions/workflows/ci.yaml)
[![license](https://img.shields.io/github/license/PlanVX/aweme?style=for-the-badge)](https://github.com/PlanVX/aweme/blob/main/LICENSE)
[![Codecov](https://img.shields.io/codecov/c/github/PlanVX/aweme?style=for-the-badge)](https://codecov.io/gh/PlanVX/aweme)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/planvx/aweme?style=for-the-badge)](https://hub.docker.com/r/planvx/aweme)

## 项目结构

```
.
├── .github
├── cmd
│   └── aweme/main.go
├── configs
│   └── config.yml
├── docker-compose.yml
├── docs
├── pkg
├── readme.md
├── renovate.json
└── scripts
    ├── codegen.sh
    ├── schema.sql
    ├── .env
    └── Dockerfile
```

- .github/workflows: 该目录包含了 GitHub Actions 的配置文件。
- cmd: 该目录包含了项目的入口包，即 main() 函数所在。
- configs: 该目录包含了项目的配置文件，例如：数据库配置、日志配置等。
- docker-compose.yml: 该文件用于在 Docker 环境下运行项目时所需的配置。
- docs: 该目录包含项目的 API 文档, 可以使用 codegen.sh 脚本生成。
- internal: 该目录包含项目的内部代码逻辑，例如：API 定义、数据访问层、逻辑层等。
- scripts: 该目录包含项目的脚本文件，例如：代码生成脚本、数据库脚本等。
- scripts/.env 文件包含了作为环境变量的配置信息，请按需填写，例如：数据库连接信息、密钥等。

## kubernetes 部署

以下操作在默认的 namespace 下进行，如果需要部署到其他 namespace，请在命令中指定。

1. 创建配置

   ```bash
   kubectl create cm aweme-config --from-file=aweme-app-config.yaml=configs/config.yml --from-file=aweme-app-schema.sql=scripts/schema.sql
   ```

2. 创建 secret，其中包括了数据库连接信息、密钥等

   ```bash
   kubectl -f deploy/secret.yaml
   ```

3. 部署 MySQL 和 redis 及相关 headless service

   ```bash
   kubectl -f deploy/mysql.yaml
   kubectl -f deploy/redis.yaml
   ```

4. 部署应用及相关 service

   ```bash
   kubectl -f deploy/aweme.yaml
   ```

5. （可选）配置 HPA 水平自动伸缩

   ```bash
   kubectl -f deploy/aweme-hpa.yaml
   ```

## 运行

项目运行可参考以下步骤：

1. 在项目目录下，运行 scripts/codegen.sh 脚本来生成项目所需的代码，包括测试代码和 API 文档，（可选）在运行该脚本时，指定 MOCKERY 环境变量，以告诉脚本使用 Mockery 工具生成相应的 Mock 代码用于逻辑层测试。
2. 根据实际情况修改 scripts 下的 .env 文件，并在其中配置环境变量，例如：数据库连接信息、密钥等。
3. 运行 docker-compose up 命令来启动项目。该命令会根据 docker-compose.yml
   文件中的配置来构建和启动项目的容器。
