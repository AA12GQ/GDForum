package limiter

import (
	"GDForum/pkg/config"
	"GDForum/pkg/logger"
	"GDForum/pkg/redis"
	"github.com/gin-gonic/gin"
	"strings"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)


// GetKeyIP 获取 Limitor 的 Key，IP
func GetKeyIP(c *gin.Context)string{
	return c.ClientIP()
}

// GetKeyRouteWithIP Limitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context)string{
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

// CheckRate 检测请求是否超额
func CheckRate(c *gin.Context,key string,formatted string)(limiterlib.Context, error){
	var context limiterlib.Context
	rate,err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context,err
	}
	store,err := sredis.NewStoreWithOptions(redis.Redis.Client,limiterlib.StoreOptions{
		Prefix:config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context,err
	}
	limiterObj := limiterlib.New(store,rate)
	return limiterObj.Get(c,key)
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string)string{
	routeName = strings.ReplaceAll(routeName,"/","-")
	routeName = strings.ReplaceAll(routeName,":","_")
	return routeName
}
