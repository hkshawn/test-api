package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// App  这只是个返回的结构,不是表结构
type App struct {
	ClientToken      string      `json:"client_token"`       // 客户端 token(客户端读取信息用)
	GatewayRules     string      `json:"gateway_rules"`      // 监听端口,例如: 11:11;22:22;456:456-789
	ServerAddr       string      `json:"server_addr"`        // 服务器的外网地址，例如：192.168.1.1:8888(服务器真实 IP)
	ServerListenAddr string      `json:"server_listen_addr"` // 服务器的监听地址，例如：0.0.0.0:8888
	ClientListenIP   string      `json:"client_listen_ip"`   // 客户端的本地监听IP，例如：127.0.0.2
	Proxies          []Proxy     `json:"proxies"`
	DummyNodes       []DummyNode `json:"dummy_nodes"`
}

func (o App) TableName() string {
	return "app"
}

// AppProxyDummyNode 中间件表  举例client_token为 111 的用户拥有 tx1 分组的 proxy 和 dn1 的 dummy_node
type AppProxyDummyNode struct {
	ClientToken  string `json:"client_token"` // client_token
	ServerToken  string `json:"server_token"`
	ProxyTag     string `json:"proxy_tag"`      // proxy_tag
	DummyNodeTag string `json:"dummy_node_tag"` //dummy_node_tag
}

func (o AppProxyDummyNode) TableName() string {
	return "app_proxy_dummy_node"
}

type Proxy struct {
	ProxyToken string `json:"proxy_token"`
	Network    string `json:"network"`
	Addr       string `json:"addr"`
}

func (o Proxy) TableName() string {
	return "proxy"
}

// DummyNode 虚拟节点表
type DummyNode struct {
	DummyNodeTag string `gorm:"dummy_node_tag"` // 指定外键
	Network      string `json:"network"`        // 节点的类型，tcp/kcp/quic
	Addr         string `json:"addr"`           // 节点的外网地址
}

func (o DummyNode) TableName() string {
	return "dummy_node"
}

func FindApp(CT string) App {
	var app App
	if e := DB.Model(App{}).Where("client_token=?", CT).Find(&app).Error; e != nil {
		panic(e)
	}
	return app
}

func FindProxy(PT string) []Proxy {
	var app []Proxy
	if e := DB.Where("proxy_tag=?", PT).Find(&app).Error; e != nil {
		panic(e)
	}
	return app
}

func FindNode(DT string) []DummyNode {
	var app []DummyNode
	if e := DB.Where("dummy_node_tag=?", DT).Find(&app).Error; e != nil {
		panic(e)
	}
	return app
}

func FindMidl(CT string) AppProxyDummyNode {

	var app AppProxyDummyNode
	//get from app table
	if e := DB.Where("client_token=?", CT).Find(&app).Error; e != nil {
		panic(e)
	}
	return app
}
