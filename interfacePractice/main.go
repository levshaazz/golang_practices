package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    string
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e

	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id does not exist")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

/*		main для проверки интерфейса storage
func main() {
	var s storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T.\n", s)

	s = newMemoryStorage()

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T.\n", s)
}
*/

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}

	return e, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Printf("удаление пользователя с id: %d прошло успешно\n", id)
	return nil
}

// main, чтобы тестировать интерфейс storage и dumbStorage
/*
func main() {
	var s storage

	fmt.Println("s:", s)
	fmt.Printf("type of s: %T.\n\n", s)

	s = newMemoryStorage()

	fmt.Println("s:", s)
	fmt.Printf("type of s: %T.\n\n", s)

	s = newDumbStorage()

	fmt.Println("s:", s)
	fmt.Printf("type of s: %T.\n\n", s)

	s = nil

	fmt.Println("s:", s)
	fmt.Printf("type of s: %T.\n\n", s)
}
*/

// main, чтобы тестировать spawnEmployees
func main() {
	nms := newMemoryStorage()
	nds := newDumbStorage()

	spawnEmployees(nms)
	fmt.Println(nms.get(5))

	spawnEmployees(nds)
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}
