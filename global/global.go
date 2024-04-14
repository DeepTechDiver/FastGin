package global

import (
	"fast-gin/global/config"
	"fast-gin/global/database/mysql"
	RedisDbFun "fast-gin/global/database/redis"
	log "fast-gin/global/logrus"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	Logger = log.ReturnsInstance()
	RedisDb = RedisDbFun.ReturnsInstance()
	Db = mysql.ReturnsInstance()
	Config = config.ReturnsInstance()

}

var (
	Logger  *logrus.Logger
	Config  *config.Info
	Db      *gorm.DB
	RedisDb *redis.Client
)
