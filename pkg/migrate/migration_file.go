package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator,*sql.DB)

var migrationFiles []MigrationFile

// MigrationFile 代表着单个迁移文件
type MigrationFile struct {
	Up 			migrationFunc
	Down 		migrationFunc
	FileName 	string
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string,up migrationFunc,down migrationFunc){
	migrationFiles = append(migrationFiles,MigrationFile{
		Up:       up,
		Down:     down,
		FileName:	name,
	})
}
