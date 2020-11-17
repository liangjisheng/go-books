package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/oschwald/geoip2-golang"
)

// Response ...
type Response struct {
	Country   string
	Province  string
	City      string
	ISP       string
	Latitude  float64
	Longitude float64
	TimeZone  string
}

// IP2addr ...
type IP2addr struct {
	db *geoip2.Reader
}

// Agrs  ...
type Agrs struct {
	IPString string
}

// Address ...
func (t *IP2addr) Address(agr *Agrs, res *Response) error {
	netIP := net.ParseIP(agr.IPString)
	// 调用开源geoIp 数据库查询ip地址
	record, err := t.db.City(netIP)
	res.City = record.City.Names["zh-CN"]
	res.Province = record.Subdivisions[0].Names["zh-CN"]
	res.Country = record.Country.Names["zh-CN"]
	res.Latitude = record.Location.Latitude
	res.Longitude = record.Location.Longitude
	res.TimeZone = record.Location.TimeZone
	return err
}

func main() {
	// 加载geoIp数据库
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	// 初始化jsonRPC
	ip2addr := &IP2addr{db}
	// 注册
	rpc.Register(ip2addr)
	// 绑定端口
	address := ":3344"
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("json rpc is listening", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
