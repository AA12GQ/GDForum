package validators

import (
	"GDForum/pkg/database"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

//注册自定义表单验证规则
func init(){
	govalidator.AddCustomRule("not_exists", func(field string, rule string,
		message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)

		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		var count int64
		query.Count(&count)

		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		return nil
	})

}
