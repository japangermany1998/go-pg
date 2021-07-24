package main

import (
	"go-pg/config"
	"go-pg/controller"
	"go-pg/router"

	"github.com/kataras/iris/v12"
)

func main() {
	db := config.ConnectDatabase() //Kết nối database
	defer db.Close()               //Đóng database khi kết thúc chương trình
	app := iris.New()              //Sử dụng framework iris
	controller.DB = db

	router.AllRoutes(app)
	app.Run(iris.Addr(":8080"))

}
