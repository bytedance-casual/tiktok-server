package sources

import (
	"tiktok-server/internal/conf"
	"tiktok-server/internal/model"
)

var Context = &dataSourceContext{}

type IDataSource interface {
	Setup(config *model.Config) // 数据库初始化接口
	Close()                      // 关闭连接接口
	ConfigAvailable(config *model.Config) bool
}

type dataSourceContext struct {
	sources []IDataSource
}

func (ctx *dataSourceContext) RegisterSource(source IDataSource) {
	ctx.sources = append(ctx.sources, source)
}

func (ctx *dataSourceContext) Setup() {
	global := conf.Config
	for _, source := range ctx.sources {
		if source.ConfigAvailable(global) {
			source.Setup(global)
		}
	}
}
