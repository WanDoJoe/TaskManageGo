package main

import (
	_ "taskmanage/routers"
	"github.com/astaxie/beego"
	"taskmanage/models"
)

func main() {
	models.Init();
	beego.Run()
}

