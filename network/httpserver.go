package network

import (
	"fmt"
	"log"
	"net/http"

	owlconfig "github.com/xssed/owlcache/config"
	//owlhttp "github.com/xssed/owlcache/network/httpserver"
)

func stratHTTP() {

	addr := owlconfig.OwlConfigModel.Host + ":" + owlconfig.OwlConfigModel.Httpport
	//默认信息
	http.HandleFunc("/", IndexPage) //设置访问的路由
	//单机数据执行信息
	http.HandleFunc("/data/", Exe) //设置访问的路由
	//群组数据执行信息
	http.HandleFunc("/group_data/", GroupExe) //设置访问的路由
	//设置服务器集群
	http.HandleFunc("/server_group/", ServerGroup) //设置服务器集群

	//监听设置
	err := http.ListenAndServe(addr, nil) //设置监听的端口
	if err != nil {
		fmt.Println("Error starting HTTP server.")
		log.Fatal("ListenAndServe: ", err)
	}

}

//单机数据执行信息
func Exe(w http.ResponseWriter, r *http.Request) {

	owlhandler := NewOwlHandler()
	owlhandler.owlrequest.HTTPReceive(w, r)
	owlhandler.HTTPHandle(w, r) //执行数据
	resstr := owlhandler.owlresponse.ConvertToString()
	fmt.Fprintf(w, resstr) //输出到客户端的信息

}

//群组数据执行信息
func GroupExe(w http.ResponseWriter, r *http.Request) {

	owlhandler := NewOwlHandler()
	owlhandler.owlrequest.HTTPReceive(w, r)
	owlhandler.HTTPGroupDataHandle(w, r) //执行数据
	resstr := owlhandler.owlresponse.ConvertToString()
	fmt.Fprintf(w, resstr) //输出到客户端的信息

}

//设置服务器集群
func ServerGroup(w http.ResponseWriter, r *http.Request) {

	owlhandler := NewOwlServerGroupHandler()
	//owlhandler.owlrequest.HTTPReceive(w, r)
	owlhandler.HTTPServerGroupHandle(w, r) //执行数据
	//	resstr := owlhandler.owlresponse.ConvertToString()
	//	fmt.Fprintf(w, resstr) //输出到客户端的信息

}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to use Owlcache.") //输出到客户端的信息
}
