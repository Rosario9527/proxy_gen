package xtun2socks

import (
	// _ "github.com/xjasonlyu/tun2socks/v2/dns" // 引入这一句，会导致v2ray，解析失败
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"go.uber.org/automaxprocs/maxprocs"
)

type Options struct {
	MTU      int
	Proxy    string
	LogLevel string
	Device   string
}

func Start(opt *Options) {
	ch := make(chan struct{})
	go func() {
		close(ch)
		key := new(engine.Key)
		key.MTU = opt.MTU
		key.Device = opt.Device
		key.LogLevel = opt.LogLevel
		key.Proxy = opt.Proxy
		maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {}))
		engine.Insert(key)
		engine.Start()
	}()
	<-ch
}

func Stop() {
	engine.Stop()
}
