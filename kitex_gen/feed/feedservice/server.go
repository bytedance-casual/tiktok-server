// Code generated by Kitex v0.4.4. DO NOT EDIT.
package feedservice

import (
	server "github.com/cloudwego/kitex/server"
	feed "tiktok-server/kitex_gen/feed"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler feed.FeedService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
