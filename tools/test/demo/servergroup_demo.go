package main

import (
	"fmt"

	"github.com/xssed/owlcache/group"
)

type OwlServerGroupRequest struct {
	//请求命令
	Cmd string
	//地址字符串
	Address string
	//链接密码
	Pass string
	//token
	Token string
}

var servergroup *group.Servergroup

func main() {

	servergroup = group.NewServergroup()
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.1", "1111111111111", ""})
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.2", "1111111111111", ""})
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.3", "1111111111111", ""})
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.4", "1111111111111", ""})
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.5", "1111111111111", ""})
	servergroup.Add(OwlServerGroupRequest{"", "http://192.168.0.6", "1111111111111", ""})
	servergroup.AddAt(2, OwlServerGroupRequest{"", "http://192.168.0.7", "1111111111111", ""})
	//	fmt.Println(servergroup.AddAt(2, OwlServerGroupRequest{"", "http://192.168.0.8", "1111111111111", ""}))
	//	fmt.Println(servergroup.AddAt(222, OwlServerGroupRequest{"", "http://192.168.0.8", "1111111111111", ""}))
	//	fmt.Println(servergroup.ToSliceString())
	//	servergroup.RemoveFirst()
	//	fmt.Println(servergroup.ToSliceString())
	//	servergroup.RemoveLast()
	//	fmt.Println(servergroup.ToSliceString())
	//	servergroup.RemoveAt(1)
	//	fmt.Println(servergroup.ToSliceString())
	//	servergroup.Clear()
	//	fmt.Println(servergroup.ToSliceString())
	//	fmt.Println(servergroup.GetFirst())
	//	fmt.Println(servergroup.GetLast())
	//	fmt.Println(servergroup.GetAt(2))
	//	fmt.Println(servergroup.GetAt(2222))
	//	fmt.Println(servergroup.GetRange(2, 3))
	//	fmt.Println(servergroup.Count())
	resat := 0
	resbool := false
	list := servergroup.Values()
	fmt.Println(list)
	for k := range list {
		val, ok := list[k].(OwlServerGroupRequest)
		if ok {
			if val.Address == "http://192.168.0.14" {
				resat = k
				resbool = true
			}
		}
	}
	fmt.Println(resat)
	fmt.Println(resbool)

	//fmt.Println(servergroup.Exists("192.168.1.10"))

}
