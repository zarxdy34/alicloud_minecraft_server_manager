// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/zarxdy34/alicloud_minecraft_server_manager/src/biz"
)

func Init() {
	biz.InitConfig()
}

func main() {
	h := server.Default()

	Init()
	register(h)
	h.Spin()
}
