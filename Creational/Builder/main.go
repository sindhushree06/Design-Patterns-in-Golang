package builder

import "fmt"

type Student struct {
	id   int    `bson:"_id"`
	name string `bson:"name"`
}

type StudentBuilder struct {
	student Student
}

func NewStudentBuilder() *StudentBuilder {
	return &StudentBuilder{}
}

func (s *StudentBuilder) CreateStudentID(ID int) *StudentBuilder {
	s.student.id = ID
	return s
}

func (s *StudentBuilder) CreateStudentName(name string) *StudentBuilder {
	s.student.name = name
	return s
}

func (s *StudentBuilder) build() Student {
	return s.student
}

func main() {
	obj := NewStudentBuilder().
		CreateStudentID(123).
		CreateStudentName("sindhu").
		build()

	fmt.Println(obj)

}
