package xtun2socks

import (
	_ "github.com/xjasonlyu/tun2socks/v2/dns"
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"go.uber.org/automaxprocs/maxprocs"
)

type Options = engine.Key

func Start(opt *Options) {
	ch := make(chan struct{})
	go func() {
		close(ch)
		maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {}))
		engine.Insert(opt)
		engine.Start()
	}()
	<-ch
}

func Stop() {
	engine.Stop()
}
