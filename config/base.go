package config

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Rdb redis连接
var Rdb *redis.Client

// Db  数据库
var Db *gorm.DB
