package model

type Student struct {
	Id   int
	Name string

	TeacherID int
}

type Teacher struct {
	Id   int
	Name string

	// has many
	Student []Student
}
