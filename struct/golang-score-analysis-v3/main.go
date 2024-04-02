package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func (s *School) GetGrades() []int {
	return s.Grades
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	var min int
	var max int

	min = s.Grades[0]
	for _, grade := range s.Grades {
		if grade < min {
			min = grade
		}
	}

	max = s.Grades[0]
	for _, grade := range s.Grades {
		if grade > max {
			max = grade
		}
	}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}

	average := float64(sum) / float64(len(s.Grades))

	return average, min, max // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}
