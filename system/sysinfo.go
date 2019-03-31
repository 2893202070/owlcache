package system

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const (
	VERSION      string = "0.1.2"
	VERSION_DATE string = "2018-01-30"
)

//程序启动欢迎信息
func DosSayHello() {

	fmt.Println("\nWelcome to use owlcache. Version:" + VERSION + "\nIf you have any questions,Please contact us: xsser@xsser.cc \nProject Home:https://github.com/xssed/owlcache")
	fmt.Println(`                _                _          `)
	fmt.Println(`   _____      _| | ___ __ _  ___| |__   ___ `)
	fmt.Println(`  / _ \ \ /\ / / |/ __/ _' |/ __| '_ \ / _ \`)
	fmt.Println(` | (_) \ V  V /| | (_| (_| | (__| | | |  __/`)
	fmt.Println(`  \___/ \_/\_/ |_|\___\__,_|\___|_| |_|\___|`)
	fmt.Println(`                                            `)

}

//http服务欢迎页
func HttpSayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<style type='text/css'>*{ padding: 0; margin: 0; } div{ padding: 4px 48px;} a{color:#2E5CD5;cursor: pointer;text-decoration: none} a:hover{text-decoration:underline; } body{ background: #fff; font-family: 'Century Gothic','Microsoft yahei'; color: #333;font-size:18px;} h1{ font-size: 100px; font-weight: normal; margin-bottom: 12px; } p{ line-height: 1.6em; font-size: 42px }</style><div style='padding: 24px 48px;'><h1>:)</h1><p>Welcome to use owlcache. Version:"+VERSION+"<br/><span style='font-size:25px'>If you have any questions,Please contact us: <a href=\"mailto:xsser@xsser.cc\">xsser@xsser.cc</a><br>Project Home : <a href=\"https://github.com/xssed/owlcache\" target=\"_blank\">https://github.com/xssed/owlcache</a></span></p><div>")
}

//捕获程序正常退出操作 ctrl+c
func OnExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	log.Println("owlcache is stoped") //日志记录
	fmt.Println("owlcache  is stoped \nBye!")
}
