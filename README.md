# kss-cli

## target

创建常用的一些列操作的 cli 工具，通过命令方式能快速实现自己的的需求。

## list

[ ] 本地创建一个 git 仓库，并且将远程仓库和本地仓库关联起来。  

## 使用 go 的原因是方便编译成跨平台的可执行文件，方便部署，最主要是编译的速度快

kss-cli/
├── cmd/
│   └── app/             # 主应用程序或二进制文件
│       └── main.go
├── pkg/                 # 公共库，对外暴露的代码
├── internal/            # 私有库，只供当前项目使用
├── api/                 # API 定义（protobuf、swagger 等）
├── configs/             # 配置文件
├── scripts/             # 脚本文件（如部署、构建、CI 等）
├── deployments/         # 部署文件（如 Docker、K8s、Helm 等）
├── web/                 # 前端资源（如静态文件或模板）
├── docs/                # 文档
├── test/                # 测试文件或数据
├── Makefile             # Makefile 构建
└── README.md

## 命令

``` bash
go build -o kss-cli./cmd/app/main.go
```
