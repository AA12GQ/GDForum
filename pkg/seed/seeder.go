package seed

import (
	"GDForum/pkg/console"
	"GDForum/pkg/database"
	"gorm.io/gorm"
)

var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders 数组中
func Add(name string,fn SeederFunc){
	seeders = append(seeders,Seeder{
		Func: fn,
		Name: name,
	})
}

// SetRunOrder 设置『按顺序执行的 Seeder 数组』
func SetRunOrder(names []string){
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder{
	for _,sdr := range seeders{
		if name == sdr.Name{
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有 Seeder
func RunAll(){

	executed := make(map[string]string)
	for _,name := range orderedSeederNames{
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0{
			console.Warning("Running Odered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	//再运行剩下的
	for _,sdr := range seeders{
		//过滤已运行
		if _,ok := executed[sdr.Name]; !ok{
			console.Warning("Running Seeder:" + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

func RunSeeder(name string){
	for _,sdr := range seeders{
		if sdr.Name == name{
			sdr.Func(database.DB)
			break
		}
	}
}