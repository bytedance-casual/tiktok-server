// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishservice

import (
	server "github.com/cloudwego/kitex/server"
	publish "tiktok-server/kitex_gen/publish"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler publish.PublishService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
