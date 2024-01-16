// Code generated by Kitex v0.7.3. DO NOT EDIT.
package dionysusbasic

import (
	server "github.com/cloudwego/kitex/server"
	basic "github.com/supermicah/dionysus/basic/kitex_gen/dionysus/basic"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler basic.DionysusBasic, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
