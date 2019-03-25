package httpclient

import (
	"fmt"
	owllog "log"
	"net/http"
	"os"
	"time"

	owlconfig "github.com/xssed/owlcache/config"
)

//定义HTTP客户端结构
type OwlClient struct {
	OwlTransport     *http.Transport
	HCRequestTimeout *time.Duration
}

//创建一个HTTP客户端
func NewOwlClient() *OwlClient {
	//创建Transport()
	owltransport := NewOwlTransport()
	//从配置中取出集群互相通信时的请求超时时间
	hcrequesttimeout, err := time.ParseDuration(owlconfig.OwlConfigModel.HttpClientRequestTimeout + "s")
	if err != nil {
		//强制异常，退出
		owllog.Println("HttpClientRequestTimeout Parse error：" + err.Error()) //日志记录
		fmt.Println("HttpClientRequestTimeout Parse error：" + err.Error())
		os.Exit(0)
	}
	owlhttpclient := &OwlClient{OwlTransport: owltransport, HCRequestTimeout: &hcrequesttimeout}

	return owlhttpclient
}

//发送GET请求
func (c *OwlClient) Get(address string, value string) {
	owlclient := NewOwlHttpClient(c.OwlTransport)
	fmt.Println(owlclient)
	owlclient.Get("httpbin.org/get")
	owlclient.SetTimeout(*c.HCRequestTimeout * time.Second)
	owlclient.Query.Add("key", "value")
	res, err := owlclient.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.String())

	owlclient.Claer()
}
