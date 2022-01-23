package seeders

import (
	"GDForum/database/factories"
	"GDForum/pkg/console"
	"GDForum/pkg/logger"
	"GDForum/pkg/seed"
	"gorm.io/gorm"
	"fmt"
)

func init(){

	//添加Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factories.MakeUsers(10)

		//批量创建用户
		result := db.Table("users").Create(&users)
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded",
			result.Statement.Table, result.RowsAffected))
	})
}
