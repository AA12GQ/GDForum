package bootstrap

import (
	"GDForum/pkg/config"
	"GDForum/pkg/redis"
	"fmt"
)

//初始化redis
func SetupRedis(){
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v",config.GetString("redis.host"),config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis_password"),
		config.GetInt("redis.port"),
		)
}
