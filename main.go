package main

import (
	"demo/dao"
	"demo/models"
	"demo/routers"
	"demo/utils"
)

func main() {
	config := utils.InitConfig("config")

	err := dao.InitMySQL(config)
	if err != nil {
		panic(err)
	}

	defer dao.Close()

	dao.DB.AutoMigrate(&models.Qq{})
	dao.DB.AutoMigrate(&models.Weibo{})

	r := routers.SetupRouter()
	r.Run(":"+config["listen_port"])
}
