package main

import (
	"awesomeProject/model"
	"awesomeProject/router"
	"log"
)

func main() {
	//初始数据库
	model.DBInit()
	//初始路由
	engine := router.CreateRouter()
	//启动
	err := engine.RunTLS(":3333", "https/gelmezon_cn_integrated.crt", "https/gelmezon_cn.key")
	if err != nil {
		log.Default().Println(err.Error())
	}
	defer model.DbClient.Close()
}
