package main

import (
	"com.justin.k8s.api/common"
	"com.justin.k8s.api/srv-article/routers"
)

func main() {
	common.Init()
	defer common.Release()

	routers.InitRouter()
}
