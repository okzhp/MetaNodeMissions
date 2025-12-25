package main

import (
	"auctionBackEnd/config"
	"auctionBackEnd/ether"
	"auctionBackEnd/router"
	"log"
)

func main() {
	//初始化数据库
	db := config.InitDB()

	//单独启动一个线程监听链上事务
	go ether.ListeningChainEvent(db)

	//设置路由
	engine := router.SetRouter(db)

	//启动服务
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalln("服务启动失败:", err)
	}
}
