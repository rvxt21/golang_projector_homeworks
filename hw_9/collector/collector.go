package collector

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Storage struct {
	m           sync.Mutex
	lastID      int
	allStudents map[int]Student
}

func NewStorage() *Storage {
	return &Storage{
		allStudents: make(map[int]Student),
	}
}

func (s *Storage) GetAllStudents() []Student {
	s.m.Lock()
	defer s.m.Unlock()

	var students = make([]Student, 0, len(s.allStudents))

	for _, student := range s.allStudents {
		students = append(students, student)
	}

	sort.Slice(students, func(i, j int) bool { return students[i].ID < students[j].ID })

	return students
}

func (s *Storage) CreateOneStudent(st Student) int {
	s.m.Lock()
	defer s.m.Unlock()

	fmt.Println("Trying to create student")

	st.ID = s.lastID + 1
	s.allStudents[st.ID] = st
	s.lastID++
	st.UpdatedAt = time.Now()

	fmt.Printf("Last ID:%d\n", s.lastID)
	return st.ID
}

func (s *Storage) FindStudentById(id int) (Student, bool) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, st := range s.allStudents {
		if st.ID == id {
			return st, true
		}
	}
	return Student{}, false
}
