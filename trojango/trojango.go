package trojango

import (
	"github.com/p4gefau1t/trojan-go/log"
	"github.com/p4gefau1t/trojan-go/proxy"
)

var (
	proxy_ *proxy.Proxy
)

func Start(jsonCfg string) {
	ch := make(chan struct{})
	go func() {
		close(ch)
		var err error
		proxy_, err = proxy.NewProxyFromConfigData([]byte(jsonCfg), true)
		if err != nil {
			log.Fatal(err)
		}
		err = proxy_.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-ch
}

func Stop() {
	if proxy_ != nil {
		proxy_.Close()
		proxy_ = nil
	}
}
