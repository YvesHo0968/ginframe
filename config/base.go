package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Rdb redis连接
var Rdb *redis.Client

var RdbCtx = context.Background()

// Db  数据库
var Db *gorm.DB

// Log 日志
var Log zerolog.Logger

// Flag flag
var Flag FlagConfig

// Viper 配置数据
var Viper viperConfig

func Init() {
	// Flag
	//InitFlag()

	// 配置文件
	InitConfigFile()

	// 启动日志
	InitLog()

	// 启动redis
	InitRedis()

	// 启动数据库
	InitDb()

	Log.Debug().Msg("config ini")
}
