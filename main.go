package main

import (
	_ "github.com/perseus/go-serialpi/routers"
	"github.com/astaxie/beego"
	"github.com/perseus/go-serialpi/controllers"
)

func main() {

	beego.Router("/up", &controllers.MainController{} , "get:Up;post:Up")
	beego.Router("/down", &controllers.MainController{} , "get:Down;post:Down")
	beego.Router("/left", &controllers.MainController{} , "get:Left;post:Left")
	beego.Router("/right", &controllers.MainController{} , "get:Right;post:Right")
	beego.Router("/", &controllers.MainController{})
	beego.Run()

}

