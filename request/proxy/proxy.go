package proxy

import (
	"fmt"
	"math/rand"
	"time"
)

type ProxyConfig struct {
	Id         int
	Ip         string
	Port       int
	Location   string
	Safe       bool
	Schema     string
	VerifyTime string
	Cnt        int
	Delayed    time.Duration
}

type ProxyContainer interface {
	Init()
	One() *ProxyConfig
	Len() int
	DeleteProxy(i int)
}

type Proxy interface {
	NewProxyImpl(dsn string) (ProxyContainer, error)
}

var proxies = make(map[string]Proxy)

func Register(name string, proxy Proxy) {
	if proxy == nil {
		panic("proxy: Register proxy is nil")
	}
	if _, ok := proxies[name]; ok {
		panic("proxy: Register called twice for adapter " + name)
	}
	proxies[name] = proxy
}

func NewProxy(proxy_name, dsn string) (ProxyContainer, error) {
	proxy, ok := proxies[proxy_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown proxy_name %q", proxy_name)
	}
	return proxy.NewProxyImpl(dsn)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
