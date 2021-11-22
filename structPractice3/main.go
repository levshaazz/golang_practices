package main

import "fmt"

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Employee: %s,\nSex: %s,\nAge: %d,\nSalary: %d;", e.name, e.sex, e.age, e.salary)
}

func main() {
	employee1 := newEmployee("Roman", "Male", 27, 65000)
	employee2 := newEmployee("Alexis", "Male", 23, 99000)

	setName(&employee1, "Underpaid")

	fmt.Printf("%s\n", employee1.getInfo())
	fmt.Println("")
	fmt.Printf("%s\n", employee2.getInfo())
}

func setName(e *employee, name string) {
	e.name = name
}
