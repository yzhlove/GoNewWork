package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	Id        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Avg() float32 {
	var res float32
	for _, grade := range s.Grades {
		res += grade.Score
	}
	return res / float32(len(s.Grades))
}

type Students []Student

var (
	_students Students
	_stuMutex sync.RWMutex
)

func (ss Students) Get(id int) (*Student, error) {
	for i := range ss {
		if ss[i].Id == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with id %d not found", id)
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
