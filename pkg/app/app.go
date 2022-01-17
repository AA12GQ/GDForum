package app

import (
	"GDForum/pkg/config"
	"time"
)

func IsLocal()bool{
	return config.Get("app_env") == "local"
}

func IsProduction()bool{
	return config.Get("app_env") == "production"
}

func IsTesting()bool{
	return config.Get("app_env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}