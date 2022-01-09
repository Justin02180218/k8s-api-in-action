package main

import (
	"com.justin.k8s.api/common"
	"com.justin.k8s.api/srv-user/routers"
)

func main() {
	common.Init()
	defer common.Release()

	routers.InitRouter()
}
