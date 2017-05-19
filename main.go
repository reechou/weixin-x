package main

import (
	"github.com/reechou/weixin-x/config"
	"github.com/reechou/weixin-x/controller"
)

func main() {
	controller.NewLogic(config.NewConfig()).Run()
}
