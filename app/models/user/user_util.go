package user

import "GDForum/pkg/database"

//IsEmailExist 判断邮箱是否被注册
func IsEmailExist(eamil string)bool{
	var count int64
	database.DB.Model(User{}).Where("email = ?",eamil).Count(&count)
	return count > 0
}
// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}