package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func InitDb() {
	var err error

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/kuyo_logic_server_acc?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms", // DSN data source name
		DefaultStringSize:         256,                                                                                                         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true,                                                                                                        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// ----------------------- 连接池设置 -----------------------
	sqlDB, err := Db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetConnMaxIdleTime 设置最大空闲时间为1分钟的连接池
	sqlDB.SetConnMaxIdleTime(time.Minute)
}

func GetDB() *gorm.DB {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Println("connect db server failed.")
	}
	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
	}

	return Db
}
