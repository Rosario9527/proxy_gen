package main

import (
	"proxy_gen/trojango"
	"time"
)

var troanCfg = `
{"run_type":"client","local_addr":"127.0.0.1","local_port":40273,"remote_addr":"us.akimon.top","remote_port":443,"password":["yummy"],"log_level":2,"ssl":{"verify":true,"cert":"","cipher":"ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA:AES128-SHA:AES256-SHA:DES-CBC3-SHA","cipher_tls13":"TLS_AES_128_GCM_SHA256:TLS_CHACHA20_POLY1305_SHA256:TLS_AES_256_GCM_SHA384","sni":"","alpn":["h2","http/1.1"]},"router":{"enable":true,"bypass":["geosite:private","geoip:private","cidr:114.114.114.114/32","cidr:223.5.5.5/32","cidr:119.29.29.29/32","cidr:180.76.76.76/32","geosite:apple-cn","geosite:google-cn","geosite:tld-cn","geosite:cn","geosite:geolocation-cn","geosite:category-games@cn","geoip:cn"],"default_policy":"proxy"}}
`

func main() {
	trojango.Start(troanCfg)

	time.Sleep(time.Second * 190000)
}
