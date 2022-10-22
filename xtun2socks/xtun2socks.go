package xtun2socks

import (
	_ "github.com/xjasonlyu/tun2socks/v2/dns"
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"github.com/xjasonlyu/tun2socks/v2/log"
	"go.uber.org/automaxprocs/maxprocs"
	"gopkg.in/yaml.v3"
)

var (
	key = new(engine.Key)
)

func Start(cfg string) {
	maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {}))
	if err := yaml.Unmarshal([]byte(cfg), key); err != nil {
		log.Fatalf("Failed to unmarshal config %v", err)
	}
	engine.Insert(key)
	engine.Start()
}

func Stop() {
	engine.Stop()
}
