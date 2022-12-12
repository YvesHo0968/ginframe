package config

import (
	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Rdb redis连接
var Rdb *redis.Client

// Db  数据库
var Db *gorm.DB

// Log 日志
var Log zerolog.Logger
