package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   float64
	WorkedHours  int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	return 1 * float64(j.BaseSalary) * float64(j.WorkingMonth) / 12
}

func (s Senior) GetBonus() float64 {
	return 2*float64(s.BaseSalary)*float64(s.WorkingMonth)/12 + (s.PerformanceRate * float64(s.BaseSalary))
}

func (m Manager) GetBonus() float64 {
	return 2*float64(m.BaseSalary)*float64(m.WorkingMonth)/12 + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

type Senior struct {
	Name            string
	BaseSalary      float64
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       float64
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0
	for _, employee := range employees {
		sum += employee.GetBonus()
	}
	return sum // TODO: replace this
}

func main() {
	adi := Junior{
		Name:         "Adi",
		BaseSalary:   1000000,
		WorkedHours:  8,
		WorkingMonth: 12,
	}
	budi := Senior{
		Name:            "Budi",
		BaseSalary:      2000000,
		WorkingMonth:    12,
		PerformanceRate: 0.5,
	}
	ciko := Manager{
		Name:             "Ciko",
		BaseSalary:       3000000,
		WorkingMonth:     12,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.5,
	}
	fmt.Println(EmployeeBonus(adi))
	fmt.Println(EmployeeBonus(budi))
	fmt.Println(EmployeeBonus(ciko))
}
