package main

import (
	"github.com/serialt/genc/model"

	"gorm.io/gen"
)

// // Dynamic SQL
//
//	type Querier interface {
//		// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
//		FilterWithNameAndRole(name, role string) ([]gen.T, error)
//	}
//

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	//g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Student{}, model.Teacher{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
