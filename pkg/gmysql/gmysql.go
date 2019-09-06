package gmysql

import (
	"fmt"
	"plan/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// MysqlClient .
	MysqlClient *gorm.DB
)

// InitMySQLClient .
func InitMySQLClient() {
	var err error
	config := config.GlobalConfig.Mysql
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, "Asia%2FShanghai")
	MysqlClient, err = gorm.Open("mysql", connectString)

	if err != nil {
		panic(err)
	}
}
