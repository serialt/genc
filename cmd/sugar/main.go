package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/serialt/genc/model"
	"github.com/serialt/genc/query"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// dsn := "root:12345678@tcp(127.0.0.1:3306)/gorm_learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置创建表名时不使用复数
		},
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Student{}, &model.Teacher{})
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	// 增
	student1 := model.Student{Name: "student1"}
	student2 := model.Student{Name: "student2"}
	student3 := model.Student{Name: "student3"}
	_ = query.Student.Create(&student1, &student2, &student3)

	teacher1 := model.Teacher{Name: "teacher1"}
	_ = query.Teacher.Create(&teacher1)

	// 删
	_, _ = query.Student.Where(query.Student.Id.Eq(3)).Delete()

	// 改
	_, _ = query.Student.Where(query.Student.Id.Eq(2)).Update(query.Student.Name, "student2_new")

	// 查
	student, _ := query.Student.Where(query.Student.Id.Eq(1)).Take()
	teacher, _ := query.Teacher.Where(query.Teacher.Id.Eq(1)).Take()

	fmt.Println(student) // {1 student1 0}
	fmt.Println(teacher) // {1 teacher1 []}

	// 关联
	_ = query.Teacher.Student.Model(&teacher1).Append(&student1, &student2)
	teacher, _ = query.Teacher.Preload(query.Teacher.Student).Where(query.Teacher.Id.Eq(1)).Take()

	fmt.Println(teacher) // {1 teacher1 [{1 student1 1} {2 student2_new 1}]}

	fmt.Println(query.Student.TableName())
	fmt.Println(query.Teacher.TableName())
}
