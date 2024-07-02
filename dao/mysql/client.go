package mysql

import (
	"fmt"

	gomysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"anime-community/config"
)

var communityClient *gorm.DB

func init() {
	dsn := getDsn(config.GetServerConfig().MysqlConfig)
	if dsn == "" {
		panic("load mysql config fail.")
	}
	db, err := gorm.Open(gomysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	communityClient = db
}

func getDsn(conf *config.MysqlConfig) string {
	if conf == nil {
		return ""
	}
	username := conf.UserName
	if conf.PassWord != "" {
		username = username + ":" + conf.PassWord
	}
	dsn := fmt.Sprintf("%s@%s(%s)/%s?charset=%s",
		username,
		conf.Protocol,
		conf.Addr,
		conf.DbName,
		conf.Charset,
	)

	return dsn + "&parseTime=true&loc=Asia%2FShanghai"
}
