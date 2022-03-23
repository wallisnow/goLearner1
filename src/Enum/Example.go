package Enum

import "fmt"

const (
	Student = iota
	Teacher
	Leader
)

func Show() {
	fmt.Println(Student, Teacher, Leader)
}
