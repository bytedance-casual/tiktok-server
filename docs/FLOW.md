# 开发流程概述

## rpc

各微服务间通过 rpc 接口相互调用功能。以 `publish` 模块为例，在处理参数过程中需要获取用户信息，则需先将 `cmd/api/rpc/user.go` 拷贝至 `cmd/publish/rpc`，而后涉及到的调用逻辑如下：

```go
func getUser() {
    resp, err := rpc.User(ctx, &user.UserRequest{UserId: userId, Token: token})
    if err != nil {
        return nil, err
    }
    return resp.User, nil	
}
```