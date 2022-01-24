package databases

import (
	"fmt"
	"log"
	"time"

	// a "com.justin.k8s.api/srv-article/model"
	// u "com.justin.k8s.api/srv-user/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitMysql() (err error) {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&",
		username,
		password,
		host,
		port,
		database,
		charset)

	DB, err = gorm.Open(driverName, args)
	if err != nil {
		log.Panic("fail to connect database, err: " + err.Error())
		return
	}

	DB.SingularTable(true)
	// DB.AutoMigrate(&u.User{}, &a.Article{})
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(10 * time.Second)

	return
}

func CloseMysql() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
