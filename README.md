# tiktok-serverdocs

## 抖音项目服务端实现

> 开发人员请阅读 [[RULES]](docs/RULES.md)

### 项目进度

功能实现顺序规划：基础功能 -> 互动功能 -> 社交功能

- [ ] 视频Feed流
- [ ] 视频投稿
- [ ] 个人主页
- [ ] 喜欢列表
- [ ] 用户评论
- [ ] 关系列表
- [ ] 消息

## 项目结构

> 目录结构参考 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

- `cmd` - 项目主要的应用程序
- `configs` - 默认配置
- `docs` - 设计和用户文档
- `idl` - idl 文件夹
- `internal` - 私有应用程序代码库
- `kitex_gen` - kitex 生成代码
- `pkg` - 外部应用程序可以使用的库代码
- `test` - 外部测试应用程序和测试数据
- `web` - 静态 Web 资源

## 技术框架

- `语言`: go
- `数据库`: 阿里云RDS
- `缓存`: redis
- `服务注册`: Etcd
- `RPC框架`: 
  - Kitex
  - Kitex拓展
    - registry-etcd
    - tracer-opentracing
- `ORM框架`:
  - GORM
  - GORM拓展
    - gorm-mysql
    - gorm-opentracing
- `HTTP框架`
  - Gin
- `链路追踪`
  - opentracing
  - jaeger

### 相关文档

- [抖音方案项目说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#)
- [极简抖音App使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [各接口说明文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)