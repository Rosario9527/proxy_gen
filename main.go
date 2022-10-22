package main

import (
	"proxy_gen/freeport"
	"proxy_gen/trojango"
	"proxy_gen/v2ray"
	"proxy_gen/xtun2socks"
)

func main() {
	freeport.GetFreePort()
	trojango.Start("{}")
	v2ray.Start("", "")
	xtun2socks.Start(nil)
}
