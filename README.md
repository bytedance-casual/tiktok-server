<p align="center">
  <img src="https://avatars.githubusercontent.com/u/122664360?s=200&v=4" width="200" height="200" alt="tiktok-server">
</p>

<div align="center">

# Tiktok Document

抖音项目服务器高性能分布式优化版实现

</div>

<p align="center">
  <a href="https://github.com/bytedance-casual/tiktok-server/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/bytedance-casual/tiktok-server" alt="license">
  </a>
  <a href="https://github.com/bytedance-casual/tiktok-server/releases">
    <img src="https://img.shields.io/github/v/release/bytedance-casual/tiktok-server?color=blueviolet&include_prereleases" alt="release">
  </a>
</p>

## 抖音项目服务端实现

> 开发人员手册 [[RULES]](docs/RULES.md)

### 相关文档

- [抖音方案项目说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [各接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)

## 项目结构

> 目录结构参考 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

- `cmd` - 项目主要的应用程序
  - `api` - 微服务网关
  - `comment` - 评论模块实现
  - `favorite` - 点赞模块实现
  - `feed` - 视频流模块实现
  - `message` - 消息模块实现
  - `publish` - 视频上传模块实现
  - `relation` - 关系模块实现
  - `user` - 用户模块实现
- `configs` - 配置文件夹
- `docs` - 设计和用户文档
- `idl` - idl 文件夹
- `internal` - 私有应用程序代码库
  - `bound` 主机信息相关
  - `conf` 配置文件相关
  - `lock` 锁操作相关
  - `errren` 错误枚举
  - `middleware` 中间件相关
  - `model` 通用结构体
  - `repo` 数据库操作封装
  - `sources` 数据源操作封装
  - `tracer` 链路追踪相关
  - `utils` 工具包
- `kitex_gen` - kitex 生成代码
- `pkg` - 外部应用程序可以使用的库代码
- `test` - 外部测试应用程序和测试数据
- `web` - 静态资源目录

## 编译

- `api` - [[build.sh]](cmd/api/build.sh)
- `comment` - [[build.sh]](cmd/comment/build.sh)
- `favorite` - [[build.sh]](cmd/favorite/build.sh)
- `feed` - [[build.sh]](cmd/feed/build.sh)
- `message` - [[build.sh]](cmd/message/build.sh)
- `publish` - [[build.sh]](cmd/publish/build.sh)
- `relation` - [[build.sh]](cmd/relation/build.sh)
- `user` - [[build.sh]](cmd/user/build.sh)

## 配置项

- [SQL](configs/init.sql) `configs/init.sql`
- [config](configs/config.yml) `configs/config.yml`
- [environment](internal/conf/conf.go) `internal/conf/conf.go`
- 运行时需要提供的服务
  - MySQL
  - Redis
  - ETCD
  - Aliyun-OSS
  - FFmpeg

## 运行

- 运行各微服务下构建脚本后，可执行文件与启动脚本在对应模块下的 `output` 文件夹内

## 测试

> 测试脚本在 [./test](./test) 目录下，运行前请先跳转到对应路径

- 单模块单元测试 `bash ./test.sh <module_name>`
- 联合单元测试 `bash ./test.sh all`
- 单模块基准测试 `bash ./bench.sh <module_name>`
- 联合基准测试 `bash ./bench.sh all`

## 贡献者名单

- [IllTamer](https://github.com/IllTamer)
- [favan1](https://github.com/favan1)
- [xixiwang12138](https://github.com/xixiwang12138)
- [slgx1121](https://github.com/slgx1121)