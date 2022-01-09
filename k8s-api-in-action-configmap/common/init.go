package common

import (
	"com.justin.k8s.api/common/config"
	"com.justin.k8s.api/common/databases"
)

func Init() {
	config.InitConfig()
	databases.InitMysql()
}

func Release() {
	databases.CloseMysql()
}
