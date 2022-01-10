// Package app 应用信息
package app

import "GDForum/pkg/config"

func IsLocal()bool{
	return config.Get("app_env") == "local"
}

func IsProduction()bool{
	return config.Get("app_env") == "production"
}

func IsTesting()bool{
	return config.Get("app_env") == "testing"
}