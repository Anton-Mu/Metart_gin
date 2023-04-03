package main

import (
	"Metart/internal/dao"
	"Metart/internal/model"
	"Metart/internal/router"
)

func main() {
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	// 程序退出后关闭数据库连接
	defer dao.Close()
	// 绑定模型
	dao.SqlSession.AutoMigrate(&model.User{})
	// 注册路由
	r := router.SetRouter()
	// 启动端口为8081的项目
	r.Run(":8081")
}
