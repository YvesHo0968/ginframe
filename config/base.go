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
