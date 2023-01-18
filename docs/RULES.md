# 开发人员手册

## 版本

- go 1.13+ (默认支持错误链相关API)

## 数据库

本实现采用阿里云RDS云数据作为储存中心，其连接参数如下

- address: illtamer.com
- port: 2436
- username: bytedance
- password: Aa123456_

## 分层思想

- `controller` 负责参数校验和调用结果处理
- `service` 业务逻辑实现
- `repository` 进行持久层操作

## internal 结构

- `conf` 配置文件加载相关
- `cont` controller 层
- `db` 数据库相关配置
- `model` 模型结构体
- `repo` repository 层
- `router` 路由相关
- `service` service 层

## 代码规范

- 格式
    - 推荐使用 GoLand 开发，代码保存时自动格式化
    - 其他 IDE 使用 gofmt+goimports 对代码格式化
    - 同文件内的结构体若重复 2 个及以上字段可直接复用，除通用结构体外不同源文件内结构体不可相互引用
- 注释
    - 包中声明的每个**公共**符号（变量、常量、结构体、函数）必须注释
    - 对于实现接口的方法不需要注释，但其接口方法的定义部分必须注释
    - 对于方法的注释，应诠释代码作用、代码参数意义、代码返回错误的场景
- 命名规范
    - 使用驼峰命名法，缩略词全大写(ServerHTTP)，但当其不需要导出时全小写(xmlHTTPRequest)
    - 不允许使用拼音与意义不明的单一字符
- 控制流程
    - 多使用，合理使用设计模式
    - 参数检查与功能实现需分离，如使用
      ```go
      if param == nil {
          return nil  
      }
      doSome(param)
      ```
      而非
      ```go
      if param != nil {
          doSome(param)
      }
      ```
    - 在 if-else 逻辑中，应该先判断简单、代码量少的操作逻辑，将复杂的，大段的逻辑留在分支的最后
      ```go
      if !enable {
          doEnable()
      } else {
          loadConfiguration()
          doSetup()
          doBroadcast()
          // ......
      }
      ```
- 错误与异常处理
    - 简单异常直接使用 `errors.New` 或 `fmt.Errorf` 处理即可
    - 对于存在异常上下文的嵌套错误，**必须**使用 `fmt.Errorf`，用 `%w` 占位符将错误加入错误链中

## 单元测试

- 所有关键函数**必须**增加单元测试，测试覆盖率**必须**在 50% +，接口覆盖率必须达到 80% +
- 测试分支需要满足相互独立，全面覆盖

## 优化

建议待全部接口实现完成后进行优化。性能优化顺序：数据库连接、gorm、代码、pprof接口分析