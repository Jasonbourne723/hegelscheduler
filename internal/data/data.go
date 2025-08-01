package data

import (
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hegelscheduler/internal/config"
	"log"
)

func NewDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "sjzy_dev_user:818f83e5@tcp(common-db.dev.sjzy.local:3306)/bsi_loris?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to db: %v", err)
		return nil
	}
	return db
}

var ProviderSet = wire.NewSet(
	NewJobExecutionRepo,
	NewJobRepo,
	NewDB,
	config.NewConfig,
)
